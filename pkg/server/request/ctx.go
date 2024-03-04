package request

import (
	"context"

	"k8s.io/apiserver/pkg/authentication/user"
)

type (
	userKey struct{}
	infoKey struct{}
)

// WithUser returns a copy of parent in which the user value is set
func WithUser(parent context.Context, user user.Info) context.Context {
	return WithValue(parent, userKey{}, user)
}

// UserFrom returns the value of the user key on the ctx
func UserFrom(ctx context.Context) (user.Info, bool) {
	u, ok := ctx.Value(userKey{}).(user.Info)
	return u, ok
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent context.Context, key interface{}, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}

func InfoFrom(ctx context.Context) (*Info, bool) {
	u, ok := ctx.Value(infoKey{}).(*Info)
	return u, ok
}

func WithInfo(parent context.Context, info *Info) context.Context {
	return WithValue(parent, infoKey{}, info)
}
