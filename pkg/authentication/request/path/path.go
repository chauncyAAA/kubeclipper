package path

import (
	"fmt"
	"net/http"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
)

type Authenticator struct {
	excludePaths sets.Set[string]
	prefixes     []string
}

func NewAuthenticator(excludePaths []string) (authenticator.Request, error) {
	var prefixes []string
	paths := sets.Set[string]{}
	for _, p := range excludePaths {
		p = strings.TrimPrefix(p, "/")
		if len(p) == 0 {
			// matches "/"
			paths.Insert(p)
			continue
		}
		if strings.ContainsRune(p[:len(p)-1], '*') {
			return nil, fmt.Errorf("only trailing * allowed in %q", p)
		}
		if strings.HasSuffix(p, "*") {
			prefixes = append(prefixes, p[:len(p)-1])
		} else {
			paths.Insert(p)
		}
	}
	return &Authenticator{
		excludePaths: paths,
		prefixes:     prefixes,
	}, nil
}

func (a *Authenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	pth := strings.TrimPrefix(req.URL.Path, "/")
	if a.excludePaths.Has(pth) || a.hasPrefix(pth) {
		req.Header.Del("Authorization")
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

func (a *Authenticator) hasPrefix(pth string) bool {
	for _, prefix := range a.prefixes {
		if strings.HasPrefix(pth, prefix) {
			return true
		}
	}
	return false
}
