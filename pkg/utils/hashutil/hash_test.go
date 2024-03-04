package hashutil

import "testing"

func TestEncodeAndComparePWD(t *testing.T) {
	pwd := "testPWD"
	hashPWD, err := EncryptPassword(pwd)
	if err != nil {
		t.Errorf("EncryptPassword() err: %v", err)
	}
	if ok := ComparePassword(pwd, hashPWD); !ok {
		t.Errorf("ComparePassword() err: %v", ok)
	}
}
