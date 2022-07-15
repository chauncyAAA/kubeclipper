/*
 *
 *  * Copyright 2021 KubeClipper Authors.
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

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

type OperationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1lister.OperationLister
}

type operationInformer struct {
	factory          internal.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func NewOperationInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOperationInformer(client, resyncPeriod, indexers, nil)
}

func NewFilteredOperationInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Operations().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Operations().Watch(context.TODO(), options)
			},
		},
		&corev1.Operation{},
		resyncPeriod,
		indexers,
	)
}

func (f *operationInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOperationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *operationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.Operation{}, f.defaultInformer)
}

func (f *operationInformer) Lister() corev1lister.OperationLister {
	return corev1lister.NewOperationLister(f.Informer().GetIndexer())
}