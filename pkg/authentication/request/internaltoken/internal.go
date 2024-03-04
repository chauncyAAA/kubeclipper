package internaltoken

import (
	"errors"
	"net/http"
	"strings"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientrest"

	"k8s.io/apiserver/pkg/authentication/user"

	"k8s.io/apiserver/pkg/authentication/authenticator"
)

type Authenticator struct {
	username string
	token    string
}

func New(username, token string) *Authenticator {
	return &Authenticator{username: username, token: token}
}

var ErrInvalidToken = errors.New("invalid internal token")

func (a *Authenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	username := strings.TrimSpace(req.Header.Get(clientrest.KcUserHeader))
	token := strings.TrimSpace(req.Header.Get(clientrest.KcTokenHeader))
	if username == "" || token == "" {
		return nil, false, nil
	}

	if username == a.username && token == a.token {
		return &authenticator.Response{
			User: &user.DefaultInfo{
				Name:   a.username,
				Groups: []string{user.AllAuthenticated},
			},
		}, true, nil
	}
	return nil, false, ErrInvalidToken
}
