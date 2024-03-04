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

var _ NodesInterface = (*nodes)(nil)

type NodesGetter interface {
	Nodes() NodesInterface
}

type NodesInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.Node, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1.NodeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type nodes struct {
	client rest.Interface
}

func newNodes(c *CoreV1Client) *nodes {
	return &nodes{client: c.RESTClient()}
}

func (c *nodes) Get(ctx context.Context, name string, opts v1.GetOptions) (result *corev1.Node, err error) {
	result = &corev1.Node{}
	err = c.client.Get().
		Resource("nodes").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *nodes) List(ctx context.Context, opts v1.ListOptions) (result *corev1.NodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &corev1.NodeList{}
	err = c.client.Get().
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *nodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
