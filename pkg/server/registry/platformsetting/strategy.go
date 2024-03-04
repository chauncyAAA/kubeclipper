package platformsetting

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
	_ rest.RESTCreateStrategy = PlatformSettingStrategy{}
	_ rest.RESTUpdateStrategy = PlatformSettingStrategy{}
	_ rest.RESTDeleteStrategy = PlatformSettingStrategy{}
)

type PlatformSettingStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func NewStrategy(typer runtime.ObjectTyper) PlatformSettingStrategy {
	return PlatformSettingStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*v1.PlatformSetting)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a PlatformSetting")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *v1.PlatformSetting) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, false)
}

func MatchPlatformSetting(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (PlatformSettingStrategy) NamespaceScoped() bool {
	return false
}

func (PlatformSettingStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (PlatformSettingStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (PlatformSettingStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (PlatformSettingStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (PlatformSettingStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (PlatformSettingStrategy) Canonicalize(obj runtime.Object) {
}

func (PlatformSettingStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (s PlatformSettingStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func (s PlatformSettingStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}
