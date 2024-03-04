package backup

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
	_ rest.RESTCreateStrategy = BackupStrategy{}
	_ rest.RESTUpdateStrategy = BackupStrategy{}
	_ rest.RESTDeleteStrategy = BackupStrategy{}
)

type BackupStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s BackupStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (s BackupStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) BackupStrategy {
	return BackupStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*v1.Backup)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Backup")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *v1.Backup) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, false)
}

func MatchBackup(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (BackupStrategy) NamespaceScoped() bool {
	return false
}

func (BackupStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (BackupStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (BackupStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (BackupStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (BackupStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (BackupStrategy) Canonicalize(obj runtime.Object) {
}

func (BackupStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
