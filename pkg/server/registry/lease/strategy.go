package lease

import (
	"context"
	"fmt"

	coordinationv1 "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

var (
	_ rest.RESTCreateStrategy = LeaseStrategy{}
	_ rest.RESTUpdateStrategy = LeaseStrategy{}
	_ rest.RESTDeleteStrategy = LeaseStrategy{}
)

type LeaseStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s LeaseStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func (s LeaseStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) LeaseStrategy {
	return LeaseStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*coordinationv1.Lease)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Lease")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *coordinationv1.Lease) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

func MatchLease(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (LeaseStrategy) NamespaceScoped() bool {
	return true
}

func (LeaseStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (LeaseStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (LeaseStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (LeaseStrategy) AllowCreateOnUpdate() bool {
	return true
}

func (LeaseStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (LeaseStrategy) Canonicalize(obj runtime.Object) {
}

func (LeaseStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
