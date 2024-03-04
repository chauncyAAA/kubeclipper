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

var _ ClustersInterface = (*clusters)(nil)

type ClustersGetter interface {
	Clusters() ClustersInterface
}

type ClustersInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.Cluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1.ClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type clusters struct {
	client rest.Interface
}

func newClusters(c *CoreV1Client) *clusters {
	return &clusters{client: c.RESTClient()}
}

func (c *clusters) Get(ctx context.Context, name string, opts v1.GetOptions) (result *corev1.Cluster, err error) {
	result = &corev1.Cluster{}
	err = c.client.Get().
		Resource("clusters").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *clusters) List(ctx context.Context, opts v1.ListOptions) (result *corev1.ClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &corev1.ClusterList{}
	err = c.client.Get().
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *clusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
