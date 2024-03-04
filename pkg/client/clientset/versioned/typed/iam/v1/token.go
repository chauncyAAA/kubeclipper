package v1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientset/versioned/scheme"
	iamv1 "github.com/kubeclipper/kubeclipper/pkg/scheme/iam/v1"
)

var _ TokenInterface = (*tokens)(nil)

type TokensGetter interface {
	Tokens() TokenInterface
}

type TokenInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*iamv1.Token, error)
	List(ctx context.Context, opts v1.ListOptions) (*iamv1.TokenList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type tokens struct {
	client rest.Interface
}

func newTokens(c *IamV1Client) *tokens {
	return &tokens{client: c.RESTClient()}
}

func (c *tokens) Get(ctx context.Context, name string, opts v1.GetOptions) (result *iamv1.Token, err error) {
	result = &iamv1.Token{}
	err = c.client.Get().
		Resource("tokens").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *tokens) List(ctx context.Context, opts v1.ListOptions) (result *iamv1.TokenList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &iamv1.TokenList{}
	err = c.client.Get().
		Resource("tokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *tokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("tokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
