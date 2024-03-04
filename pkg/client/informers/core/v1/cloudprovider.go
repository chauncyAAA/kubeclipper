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

// CloudProviderInformer provides access to a shared informer and lister for cloudProviders.
type CloudProviderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1lister.CloudProviderLister
}

type cloudProviderInformer struct {
	factory          internal.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func NewCloudProviderInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudProviderInformer(client, resyncPeriod, indexers, nil)
}

func NewFilteredCloudProviderInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().CloudProviders().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().CloudProviders().Watch(context.TODO(), options)
			},
		},
		&corev1.CloudProvider{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudProviderInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudProviderInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudProviderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.CloudProvider{}, f.defaultInformer)
}

func (f *cloudProviderInformer) Lister() corev1lister.CloudProviderLister {
	return corev1lister.NewCloudProviderLister(f.Informer().GetIndexer())
}
