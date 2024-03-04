package v1

import (
	"context"
	"time"

	"github.com/kubeclipper/kubeclipper/pkg/client/clientset/versioned/scheme"

	"k8s.io/client-go/rest"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	corev1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

var _ OperationsInterface = (*operations)(nil)

type OperationsGetter interface {
	Operations() OperationsInterface
}

type OperationsInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.Operation, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1.OperationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type operations struct {
	client rest.Interface
}

func newOperations(c *CoreV1Client) *operations {
	return &operations{client: c.RESTClient()}
}

func (c *operations) Get(ctx context.Context, name string, opts v1.GetOptions) (result *corev1.Operation, err error) {
	result = &corev1.Operation{}
	err = c.client.Get().
		Resource("operations").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *operations) List(ctx context.Context, opts v1.ListOptions) (result *corev1.OperationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &corev1.OperationList{}
	err = c.client.Get().
		Resource("operations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *operations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("operations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
