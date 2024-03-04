package anonymous

import (
	"net/http"
	"strings"

	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
)

type Authenticator struct{}

func NewAuthenticator() authenticator.Request {
	return &Authenticator{}
}

func (a *Authenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	auth := strings.TrimSpace(req.Header.Get("Authorization"))
	urlToken := req.URL.Query().Get("token")
	if auth == "" && urlToken == "" {
		return &authenticator.Response{
			User: &user.DefaultInfo{
				Name:   user.Anonymous,
				UID:    "",
				Groups: []string{user.AllUnauthenticated},
			},
		}, true, nil
	}
	return nil, false, nil
}
