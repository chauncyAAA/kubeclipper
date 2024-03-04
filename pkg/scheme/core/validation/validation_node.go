package validation

import (
	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
	apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

var ValidateNodeName = apimachineryvalidation.NameIsDNSSubdomain

func ValidateNode(c *corev1.Node) field.ErrorList {
	allErrs := ValidateObjectMeta(&c.ObjectMeta, false, ValidateNodeName, field.NewPath("metadata"))
	// TODO: add other validate
	return allErrs
}
