package user

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"

	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

func NewStorage(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (rest.StandardStorage, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc: func() runtime.Object {
			return &iamv1.User{}
		},
		NewListFunc: func() runtime.Object {
			return &iamv1.UserList{}
		},
		DefaultQualifiedResource: iamv1.Resource("users"),
		KeyRootFunc:              nil,
		KeyFunc:                  nil,
		ObjectNameFunc:           nil,
		TTLFunc:                  nil,
		PredicateFunc:            MatchUser,
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
		TableConvertor:           rest.NewDefaultTableConvertor(iamv1.Resource("users")),
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
