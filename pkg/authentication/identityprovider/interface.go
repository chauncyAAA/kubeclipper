package identityprovider

import (
	"net/http"

	"github.com/kubeclipper/kubeclipper/pkg/authentication/oauth"
)

type Identity interface {
	GetUserID() string
	GetUsername() string
	GetEmail() string
}

type OAuthProvider interface {
	IdentityExchange(req *http.Request) (Identity, error)
}

type OAuthProviderFactory interface {
	Type() string
	Create(options oauth.DynamicOptions) (OAuthProvider, error)
}
