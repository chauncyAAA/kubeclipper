package v1

import (
	"context"
	"time"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientset/versioned/scheme"
	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

var _ DomainsInterface = (*domains)(nil)

type DomainsGetter interface {
	Domains() DomainsInterface
}

type DomainsInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.Domain, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1.DomainList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type domains struct {
	client rest.Interface
}

func newDomains(c *CoreV1Client) *domains {
	return &domains{client: c.RESTClient()}
}

func (c *domains) Get(ctx context.Context, name string, opts v1.GetOptions) (result *corev1.Domain, err error) {
	result = new(corev1.Domain)
	err = c.client.Get().
		Resource("domains").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *domains) List(ctx context.Context, opts v1.ListOptions) (result *corev1.DomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = new(corev1.DomainList)
	err = c.client.Get().
		Resource("domains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *domains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("domains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
