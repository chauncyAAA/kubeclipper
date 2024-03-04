package validation

import (
	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateDomain(c *corev1.Domain) field.ErrorList {
	return ValidateObjectMeta(&c.ObjectMeta, false, ValidateNodeName, field.NewPath("metadata"))
}
