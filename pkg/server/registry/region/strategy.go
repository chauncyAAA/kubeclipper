package region

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
	_ rest.RESTCreateStrategy = RegionStrategy{}
	_ rest.RESTUpdateStrategy = RegionStrategy{}
	_ rest.RESTDeleteStrategy = RegionStrategy{}
)

type RegionStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s RegionStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (s RegionStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) RegionStrategy {
	return RegionStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*v1.Region)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Region")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *v1.Region) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, false)
}

func MatchRegion(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (RegionStrategy) NamespaceScoped() bool {
	return false
}

func (RegionStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (RegionStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (RegionStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (RegionStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (RegionStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (RegionStrategy) Canonicalize(obj runtime.Object) {
}

func (RegionStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
