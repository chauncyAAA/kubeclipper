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

var _ BackupsInterface = (*backups)(nil)

type BackupsGetter interface {
	Backups() BackupsInterface
}

type BackupsInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.Backup, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1.BackupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type backups struct {
	client rest.Interface
}

func newBackups(c *CoreV1Client) *backups {
	return &backups{client: c.RESTClient()}
}

func (c *backups) Get(ctx context.Context, name string, opts v1.GetOptions) (result *corev1.Backup, err error) {
	result = &corev1.Backup{}
	err = c.client.Get().
		Resource("backups").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *backups) List(ctx context.Context, opts v1.ListOptions) (result *corev1.BackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &corev1.BackupList{}
	err = c.client.Get().
		Resource("backups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *backups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("backups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
