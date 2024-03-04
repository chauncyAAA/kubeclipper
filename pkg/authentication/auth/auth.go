package auth

import (
	"context"
	"fmt"
	"net/mail"

	v12 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"

	"github.com/kubeclipper/kubeclipper/pkg/utils/hashutil"

	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/kubeclipper/kubeclipper/pkg/query"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"

	authuser "k8s.io/apiserver/pkg/authentication/user"

	authoptions "github.com/kubeclipper/kubeclipper/pkg/authentication/options"
	"github.com/kubeclipper/kubeclipper/pkg/models/iam"
)

var (
	ErrRateLimitExceeded      = fmt.Errorf("auth rate limit exceeded")
	ErrIncorrectPassword      = fmt.Errorf("incorrect password")
	ErrAccountIsNotActive     = fmt.Errorf("account is not active")
	ErrUserNotExist           = fmt.Errorf("user not exist")
	ErrUserOrPasswordNotValid = fmt.Errorf("user or password not valid")
)

var _ PasswordAuthenticator = (*passwordAuthenticator)(nil)

type passwordAuthenticator struct {
	iamOperator iam.Operator
	authOptions *authoptions.AuthenticationOptions
}

func NewPasswordAuthenticator(operator iam.Operator, options *authoptions.AuthenticationOptions) PasswordAuthenticator {
	return &passwordAuthenticator{
		iamOperator: operator,
		authOptions: options,
	}
}

func (p *passwordAuthenticator) Authenticate(username, password string) (authuser.Info, string, error) {
	// empty username or password are not allowed
	if username == "" || password == "" {
		return nil, "", ErrUserOrPasswordNotValid
	}

	user, err := p.findUser(username)
	if err != nil {
		if !errors.IsNotFound(err) {
			return nil, "", err
		}
	}

	if user != nil && (user.Status.State == nil || *user.Status.State != v12.UserActive) {
		return nil, "", ErrAccountIsNotActive
	}

	if user != nil && user.Spec.EncryptedPassword != "" {
		if !hashutil.ComparePassword(password, user.Spec.EncryptedPassword) {
			return nil, "", ErrIncorrectPassword
		}
		return &authuser.DefaultInfo{
			Name: user.Name,
			Extra: map[string][]string{
				"phone": {user.Spec.Phone},
				"email": {user.Spec.Email},
			},
		}, "", nil
	}

	return nil, "", ErrUserNotExist
}

func (p *passwordAuthenticator) findUser(username string) (*v12.User, error) {
	if _, err := mail.ParseAddress(username); err != nil {
		return p.iamOperator.GetUserEx(context.TODO(), username, "0", false, false)
	}
	users, err := p.iamOperator.ListUsers(context.TODO(), &query.Query{
		Pagination:      query.NoPagination(),
		ResourceVersion: "0",
		FieldSelector:   fmt.Sprintf("spec.email=%s", username),
	})
	if err != nil {
		return nil, err
	}
	if len(users.Items) != 1 {
		return nil, errors.NewNotFound(v1.Resource("user"), username)
	}
	return &users.Items[0], nil
}
