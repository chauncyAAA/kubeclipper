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

// ConfigMapInformer provides access to a shared informer and lister for configMaps.
type ConfigMapInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1lister.ConfigMapLister
}

type configMapInformer struct {
	factory          internal.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func NewConfigMapInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConfigMapInformer(client, resyncPeriod, indexers, nil)
}

func NewFilteredConfigMapInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ConfigMaps().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ConfigMaps().Watch(context.TODO(), options)
			},
		},
		&corev1.ConfigMap{},
		resyncPeriod,
		indexers,
	)
}

func (f *configMapInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConfigMapInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *configMapInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.ConfigMap{}, f.defaultInformer)
}

func (f *configMapInformer) Lister() corev1lister.ConfigMapLister {
	return corev1lister.NewConfigMapLister(f.Informer().GetIndexer())
}
