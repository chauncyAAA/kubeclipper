package v1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/informers/internalinterfaces"
	"k8s.io/client-go/tools/cache"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientset"
	"github.com/kubeclipper/kubeclipper/pkg/client/internal"
	corev1lister "github.com/kubeclipper/kubeclipper/pkg/client/lister/core/v1"
	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

type RegistryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1lister.RegistryLister
}

type registryInformer struct {
	factory          internal.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func NewRegistryInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRegistryInformer(client, resyncPeriod, indexers, nil)
}

func NewFilteredRegistryInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Registries().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Registries().Watch(context.TODO(), options)
			},
		},
		&corev1.Registry{},
		resyncPeriod,
		indexers,
	)
}

func (f *registryInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRegistryInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *registryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.Registry{}, f.defaultInformer)
}

func (f *registryInformer) Lister() corev1lister.RegistryLister {
	return corev1lister.NewRegistryLister(f.Informer().GetIndexer())
}
