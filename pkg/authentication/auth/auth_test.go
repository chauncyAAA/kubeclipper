package auth

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kubeclipper/kubeclipper/pkg/utils/hashutil"

	"github.com/golang/mock/gomock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	authuser "k8s.io/apiserver/pkg/authentication/user"

	iammock "github.com/kubeclipper/kubeclipper/pkg/models/iam/mock"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

func setupIAMMock(iamMockOpera *iammock.MockOperator) {
	state := v1.UserActive
	pwd, _ := hashutil.EncryptPassword("testPWD")
	fmt.Println(pwd)
	iamMockOpera.EXPECT().ListUsers(gomock.Any(), gomock.Any()).Return(
		&v1.UserList{
			TypeMeta: metav1.TypeMeta{},
			ListMeta: metav1.ListMeta{},
			Items: []v1.User{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "testName",
					},
					Spec: v1.UserSpec{
						EncryptedPassword: pwd,
					},
					Status: v1.UserStatus{
						State: &state,
					},
				},
			},
		},
		nil).AnyTimes()
	iamMockOpera.EXPECT().GetUserEx(gomock.Any(), gomock.Any(), gomock.Eq("0"), gomock.Any(), gomock.Any()).Return(
		&v1.User{
			ObjectMeta: metav1.ObjectMeta{
				Name: "testName",
			},
			Spec: v1.UserSpec{
				EncryptedPassword: pwd,
			},
			Status: v1.UserStatus{
				State: &state,
			},
		},
		nil).AnyTimes()
}

func Test_passwordAuthenticator_Authenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	iamMockOpera := iammock.NewMockOperator(ctrl)
	setupIAMMock(iamMockOpera)

	pwdAuth := &passwordAuthenticator{
		iamOperator: iamMockOpera,
		authOptions: nil,
	}

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    authuser.Info
		want1   string
		wantErr bool
	}{
		{
			name: " empty test ",
			args: args{
				username: "",
				password: "",
			},
			want:    nil,
			want1:   "",
			wantErr: true,
		},
		{
			name: " with name pwd test ",
			args: args{
				username: "testName",
				password: "testPWD",
			},
			want: &authuser.DefaultInfo{
				Name: "testName",
				Extra: map[string][]string{
					"phone": {""},
					"email": {""},
				},
			},
			want1:   "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := pwdAuth.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf(" Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(" Authenticate() err: %v, got = %v, want %v", err, got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf(" Authenticate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
