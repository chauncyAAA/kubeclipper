package internal

import (
	"time"

	"k8s.io/apimachinery/pkg/watch"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientset"
)

type NewInformerFunc func(clientset.Interface, time.Duration) cache.SharedIndexInformer

type SharedInformerFactory interface {
	Start(stopCh <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer
}

func NewProxyWatcher(w watch.Interface, cacheSize int) watch.Interface {
	ch := make(chan watch.Event, cacheSize)
	s := watch.NewProxyWatcher(ch)
	go func() {
		defer w.Stop()
		for {
			select {
			case <-s.StopChan():
				return
			case event, ok := <-w.ResultChan():
				if !ok {
					return
				}
				e := watch.Event{
					Type: event.Type,
				}
				obj := event.Object
				if co, ok := obj.(runtime.CacheableObject); ok {
					e.Object = co.GetObject()
				} else {
					e.Object = event.Object
				}
				if s.Stopping() {
					return
				}
				ch <- e
			}
		}
	}()
	return s
}
