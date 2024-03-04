package token

import (
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

func NewStorage(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (rest.StandardStorage, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc: func() runtime.Object {
			return &v1.Token{}
		},
		NewListFunc: func() runtime.Object {
			return &v1.TokenList{}
		},
		DefaultQualifiedResource: v1.Resource("tokens"),
		KeyRootFunc:              nil,
		KeyFunc:                  nil,
		ObjectNameFunc:           nil,
		TTLFunc:                  nil,
		PredicateFunc:            MatchToken,
		EnableGarbageCollection:  false,
		DeleteCollectionWorkers:  0,
		Decorator:                nil,
		CreateStrategy:           strategy,
		BeginCreate:              nil,
		AfterCreate:              nil,
		UpdateStrategy:           strategy,
		BeginUpdate:              nil,
		AfterUpdate:              nil,
		DeleteStrategy:           strategy,
		AfterDelete:              nil,
		ReturnDeletedObject:      false,
		ShouldDeleteDuringUpdate: nil,
		TableConvertor:           rest.NewDefaultTableConvertor(v1.Resource("tokens")),
		ResetFieldsStrategy:      nil,
		Storage:                  genericregistry.DryRunnableStorage{},
		StorageVersioner:         nil,
		DestroyFunc:              nil,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return store, nil
}

//var ttlFunc = func(obj runtime.Object, existing uint64, update bool) (uint64, error) {
//	t, ok := obj.(*v1.Token)
//	if !ok {
//		return 0, fmt.Errorf("given object is not a Token")
//	}
//	if update {
//		return existing, nil
//	}
//	return uint64(*t.Spec.TTL), nil
//
//}
