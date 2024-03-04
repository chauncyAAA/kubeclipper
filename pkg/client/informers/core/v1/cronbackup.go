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

// CronBackupInformer provides access to a shared informer and lister for CronBackups.
type CronBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1lister.CronBackupLister
}

type cronBackupInformer struct {
	factory          internal.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func NewCronBackupInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterInformer(client, resyncPeriod, indexers, nil)
}

func NewFilteredCronBackupInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().CronBackups().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().CronBackups().Watch(context.TODO(), options)
			},
		},
		&corev1.CronBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *cronBackupInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCronBackupInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cronBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.CronBackup{}, f.defaultInformer)
}

func (f *cronBackupInformer) Lister() corev1lister.CronBackupLister {
	return corev1lister.NewCronBackupLister(f.Informer().GetIndexer())
}
