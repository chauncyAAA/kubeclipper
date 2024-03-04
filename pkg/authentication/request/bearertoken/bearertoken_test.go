package bearertoken

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/kubeclipper/kubeclipper/pkg/authentication/token"

	"github.com/golang/mock/gomock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"

	"github.com/kubeclipper/kubeclipper/pkg/authentication/oauth"

	"github.com/kubeclipper/kubeclipper/pkg/authentication/auth"
	authoptions "github.com/kubeclipper/kubeclipper/pkg/authentication/options"
	iammock "github.com/kubeclipper/kubeclipper/pkg/models/iam/mock"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

func authenticateRequestTestPrepare(iamMockOpera *iammock.MockOperator) {
	iamMockOpera.EXPECT().ListTokens(gomock.Any(), gomock.Any()).Return(
		&v1.TokenList{
			Items: make([]v1.Token, 1),
		}, nil).AnyTimes()
	iamMockOpera.EXPECT().GetUserEx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&v1.User{
			ObjectMeta: metav1.ObjectMeta{
				Name: "admin",
			},
			Spec: v1.UserSpec{
				Groups: make([]string, 0),
			},
		}, nil).AnyTimes()
}

func getToken() string {
	s := token.NewTokenIssuer("D9ykGOmuE3yXe35Wh3mRniGT", 0)
	u := &user.DefaultInfo{
		Name: "admin",
	}
	tokenStr, err := s.IssueTo(u, "bearer", 10*time.Hour)
	if err != nil {
		return ""
	}
	return "bearer " + tokenStr
}

func TestAuthenticator_AuthenticateRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	iamMockOpera := iammock.NewMockOperator(ctrl)
	tokenOpera := auth.NewTokenOperator(iamMockOpera, &authoptions.AuthenticationOptions{
		MaximumClockSkew: 0,
		JwtSecret:        "D9ykGOmuE3yXe35Wh3mRniGT",
		OAuthOptions: &oauth.Options{
			AccessTokenMaxAge: 10 * time.Hour,
		},
	})
	authenticateRequestTestPrepare(iamMockOpera)

	a := &Authenticator{
		auth: auth.NewTokenAuthenticator(iamMockOpera, tokenOpera),
	}

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
			name: "authenticateRequest test",
			args: args{
				req: &http.Request{
					Header: http.Header{
						"Authorization": []string{getToken()},
					},
				},
			},
			want: &authenticator.Response{
				User: &user.DefaultInfo{
					Name:   "admin",
					Groups: []string{"system:authenticated"},
				},
			},
			want1:   true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
