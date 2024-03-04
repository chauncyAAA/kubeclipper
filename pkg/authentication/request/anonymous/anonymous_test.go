package anonymous

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
)

func TestAuthenticator_AuthenticateRequest(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    *authenticator.Response
		want1   bool
		wantErr bool
	}{
		{
			name: "AuthenticateRequest empty test",
			args: args{
				req: &http.Request{
					URL: &url.URL{
						RawQuery: "token=",
					},
					Header: http.Header{
						"Authorization": []string{""},
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
			name: "AuthenticateRequest test",
			args: args{
				req: &http.Request{
					Method: "GET",
					URL: &url.URL{
						RawQuery: "token=test",
					},
					Header: http.Header{
						"Authorization": []string{""},
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
			a := &Authenticator{}
			got, got1, err := a.AuthenticateRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticateRequest() err = %v, got = %v, want %v", err, got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AuthenticateRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
