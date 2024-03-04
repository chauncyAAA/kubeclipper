package certs

import (
	"crypto"
	"crypto/x509"
	"os"
	"path/filepath"
	"testing"
)

func TestNewCaCertAndKey(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		want    *x509.Certificate
		want1   crypto.Signer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "etcd ca generate",
			args: args{cfg: Config{
				Path:         filepath.Join(pwd, "..", "dist", "pki"),
				BaseName:     "",
				CAName:       "etcd",
				CommonName:   "etcd-ca",
				Organization: []string{"etcd-ca"},
				Year:         100,
				AltNames:     AltNames{},
				Usages:       nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := NewCaCertAndKey(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCaCertAndKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
