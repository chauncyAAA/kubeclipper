package cronbackup

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
	_ rest.RESTCreateStrategy = CronBackupStrategy{}
	_ rest.RESTUpdateStrategy = CronBackupStrategy{}
	_ rest.RESTDeleteStrategy = CronBackupStrategy{}
)

type CronBackupStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (s CronBackupStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func (s CronBackupStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

func NewStrategy(typer runtime.ObjectTyper) CronBackupStrategy {
	return CronBackupStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	c, ok := obj.(*v1.CronBackup)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a CronBackup")
	}
	return c.ObjectMeta.Labels, SelectableFields(c), nil
}

func SelectableFields(obj *v1.CronBackup) fields.Set {
	return generic.AddObjectMetaFieldsSet(fields.Set{
		"spec.clusterName": obj.Spec.ClusterName,
	}, &obj.ObjectMeta, false)
}

func MatchCronBackup(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

func (CronBackupStrategy) NamespaceScoped() bool {
	return false
}

func (CronBackupStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (CronBackupStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (CronBackupStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (CronBackupStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (CronBackupStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (CronBackupStrategy) Canonicalize(obj runtime.Object) {
}

func (CronBackupStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
