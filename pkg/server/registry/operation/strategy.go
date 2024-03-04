package operation

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var (
	_ rest.RESTCreateStrategy = OperationStrategy{}
	_ rest.RESTUpdateStrategy = OperationStrategy{}
	_ rest.RESTDeleteStrategy = OperationStrategy{}
)

type OperationStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s OperationStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func (s OperationStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) OperationStrategy {
	return OperationStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*v1.Operation)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Operation")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *v1.Operation) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, false)
}

func MatchOperation(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (OperationStrategy) NamespaceScoped() bool {
	return false
}

func (OperationStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (OperationStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (OperationStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (OperationStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (OperationStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (OperationStrategy) Canonicalize(obj runtime.Object) {
}

func (OperationStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
