package wstoken

import (
	"errors"
	"net/http"

	"k8s.io/apiserver/pkg/authentication/authenticator"
)

type Authenticator struct {
	auth authenticator.Token
}

func New(auth authenticator.Token) *Authenticator {
	return &Authenticator{auth}
}

var ErrInvalidToken = errors.New("invalid bearer token")

func (a *Authenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	token := getQueryToken(req)
	if token == "" {
		return nil, false, nil
	}
	resp, ok, err := a.auth.AuthenticateToken(req.Context(), token)
	// If the token authenticator didn't error, provide a default error
	if !ok && err == nil {
		err = ErrInvalidToken
	}
	return resp, ok, err
}

// Notice: workaround for websocket auth
func getQueryToken(req *http.Request) string {
	query := req.URL.Query()
	return query.Get("token")
}
