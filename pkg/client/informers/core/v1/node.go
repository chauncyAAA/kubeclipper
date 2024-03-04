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

type NodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1lister.NodeLister
}

type nodeInformer struct {
	factory          internal.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func NewNodeInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeInformer(client, resyncPeriod, indexers, nil)
}

func NewFilteredNodeInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Nodes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Nodes().Watch(context.TODO(), options)
			},
		},
		&corev1.Node{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.Node{}, f.defaultInformer)
}

func (f *nodeInformer) Lister() corev1lister.NodeLister {
	return corev1lister.NewNodeLister(f.Informer().GetIndexer())
}
