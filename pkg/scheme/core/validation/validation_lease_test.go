package validation

import (
	"testing"

	coordinationv1 "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestValidateLease(t *testing.T) {
	lease := &coordinationv1.Lease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "invalidName++",
			Namespace: "==invalid_Namespace==",
		},
	}
	errs := ValidateLease(lease)
	if len(errs) != 2 {
		t.Errorf("unexpected list of errors: %#v", errs.ToAggregate().Error())
	}
}

func TestValidateLeaseSpec(t *testing.T) {
	holder := "holder"
	leaseDuration := int32(0)
	leaseTransitions := int32(-1)
	spec := &coordinationv1.LeaseSpec{
		HolderIdentity:       &holder,
		LeaseDurationSeconds: &leaseDuration,
		LeaseTransitions:     &leaseTransitions,
	}
	errs := ValidateLeaseSpec(spec, field.NewPath("foo"))
	if len(errs) != 2 {
		t.Errorf("unexpected list of errors: %#v", errs.ToAggregate().Error())
	}
}

func TestValidateLeaseSpecUpdate(t *testing.T) {
	holder := "holder"
	leaseDuration := int32(0)
	leaseTransitions := int32(-1)
	lease := &coordinationv1.Lease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "holder",
			Namespace: "holder-namespace",
		},
		Spec: coordinationv1.LeaseSpec{
			HolderIdentity:       &holder,
			LeaseDurationSeconds: &leaseDuration,
			LeaseTransitions:     &leaseTransitions,
		},
	}
	oldHolder := "oldHolder"
	oldLeaseDuration := int32(3)
	oldLeaseTransitions := int32(3)
	oldLease := &coordinationv1.Lease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "holder",
			Namespace: "holder-namespace",
		},
		Spec: coordinationv1.LeaseSpec{
			HolderIdentity:       &oldHolder,
			LeaseDurationSeconds: &oldLeaseDuration,
			LeaseTransitions:     &oldLeaseTransitions,
		},
	}
	errs := ValidateLeaseUpdate(lease, oldLease)
	if len(errs) != 3 {
		t.Errorf("unexpected list of errors: %#v", errs.ToAggregate().Error())
	}

	validLeaseDuration := int32(10)
	validLeaseTransitions := int32(20)
	validLease := &coordinationv1.Lease{
		ObjectMeta: oldLease.ObjectMeta,
		Spec: coordinationv1.LeaseSpec{
			HolderIdentity:       &holder,
			LeaseDurationSeconds: &validLeaseDuration,
			LeaseTransitions:     &validLeaseTransitions,
		},
	}
	validLease.ObjectMeta.ResourceVersion = "2"
	errs = ValidateLeaseUpdate(validLease, oldLease)
	if len(errs) != 0 {
		t.Errorf("unexpected list of errors for valid update: %#v", errs.ToAggregate().Error())
	}
}
