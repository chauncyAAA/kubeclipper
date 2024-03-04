package token

import (
	"reflect"
	"testing"
	"time"

	"k8s.io/apiserver/pkg/authentication/user"
)

func TestTokenIssueToAndVerify(t *testing.T) {
	s := &jwtTokenIssuer{
		name:             "kubeclipper",
		secret:           []byte("D9ykGOmuE3yXe35Wh3mRniGT"),
		maximumClockSkew: 0,
	}
	u := &user.DefaultInfo{
		Name: "admin",
	}
	token, err := s.IssueTo(u, "bearer", 10*time.Hour)
	if err != nil {
		t.Errorf("IssueTo() error = %v", err)
		return
	}
	got, _, err2 := s.Verify(token)
	if err2 != nil {
		t.Errorf("Verify() error = %v", err2)
		return
	}
	if !reflect.DeepEqual(u, got) {
		t.Errorf("Verify() got = %v, want %v", got, u)
	}
}
