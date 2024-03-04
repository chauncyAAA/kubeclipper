package globalrolebinding

import (
	"context"
	"fmt"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"

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
	_ rest.RESTCreateStrategy = GlobalRoleBindingStrategy{}
	_ rest.RESTUpdateStrategy = GlobalRoleBindingStrategy{}
	_ rest.RESTDeleteStrategy = GlobalRoleBindingStrategy{}
)

type GlobalRoleBindingStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s GlobalRoleBindingStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (s GlobalRoleBindingStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) GlobalRoleBindingStrategy {
	return GlobalRoleBindingStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*v1.GlobalRoleBinding)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a GlobalRoleBinding")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *v1.GlobalRoleBinding) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, false)
}

func MatchGlobalRoleBinding(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (GlobalRoleBindingStrategy) NamespaceScoped() bool {
	return false
}

func (GlobalRoleBindingStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (GlobalRoleBindingStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (GlobalRoleBindingStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (GlobalRoleBindingStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (GlobalRoleBindingStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (GlobalRoleBindingStrategy) Canonicalize(obj runtime.Object) {
}

func (GlobalRoleBindingStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
