package path

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
)

func TestAuthenticator_AuthenticateRequest(t *testing.T) {
	type fields struct {
		excludePaths sets.Set[string]
		prefixes     []string
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *authenticator.Response
		want1   bool
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				excludePaths: sets.New("testPath", "testPath2", "testPath3"),
				prefixes:     []string{"testPath"},
			},
			args: args{
				req: &http.Request{
					URL: &url.URL{
						Path: "testPath",
					},
					Header: http.Header{
						"Authorization": []string{"test"},
					},
				},
			},
			want: &authenticator.Response{
				User: &user.DefaultInfo{
					Name:   user.Anonymous,
					UID:    "",
					Groups: []string{user.AllUnauthenticated},
				},
			},
			want1:   true,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				excludePaths: sets.New("test Path", "testPath2", "testPath3"),
				prefixes:     []string{"testPath"},
			},
			args: args{
				req: &http.Request{
					URL: &url.URL{
						Path: "tes path",
					},
					Header: http.Header{
						"Authorization": []string{"test"},
					},
				},
			},
			want:    nil,
			want1:   false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticator{
				excludePaths: tt.fields.excludePaths,
				prefixes:     tt.fields.prefixes,
			}
			got, got1, err := a.AuthenticateRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticateRequest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AuthenticateRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
