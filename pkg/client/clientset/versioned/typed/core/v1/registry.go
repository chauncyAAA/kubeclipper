package v1

import (
	"context"
	"time"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"

	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ RegistriesInterface = (*registries)(nil)

type RegistriesGetter interface {
	Registries() RegistriesInterface
}

type RegistriesInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.Registry, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1.RegistryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type registries struct {
	client rest.Interface
}

func newRegistries(c *CoreV1Client) *registries {
	return &registries{client: c.RESTClient()}
}

func (c *registries) Get(ctx context.Context, name string, opts v1.GetOptions) (result *corev1.Registry, err error) {
	result = &corev1.Registry{}
	err = c.client.Get().
		Resource("registries").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *registries) List(ctx context.Context, opts v1.ListOptions) (result *corev1.RegistryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &corev1.RegistryList{}
	err = c.client.Get().
		Resource("registries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *registries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("registries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
