package node

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
	_ rest.RESTCreateStrategy = NodeStrategy{}
	_ rest.RESTUpdateStrategy = NodeStrategy{}
	_ rest.RESTDeleteStrategy = NodeStrategy{}
)

type NodeStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s NodeStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (s NodeStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) NodeStrategy {
	return NodeStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*v1.Node)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Node")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *v1.Node) fields.Set {
	return generic.AddObjectMetaFieldsSet(fields.Set{
		"ip": obj.Status.Ipv4DefaultIP,
	}, &obj.ObjectMeta, false)
}

func MatchNode(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (NodeStrategy) NamespaceScoped() bool {
	return false
}

func (NodeStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (NodeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (NodeStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (NodeStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (NodeStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (NodeStrategy) Canonicalize(obj runtime.Object) {
}

func (NodeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
