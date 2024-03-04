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

var _ BackupPointsInterface = (*backupPoints)(nil)

type BackupPointsGetter interface {
	BackupPoints() BackupPointsInterface
}

type BackupPointsInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.BackupPoint, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1.BackupPointList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
}

type backupPoints struct {
	client rest.Interface
}

func newBackupPoints(c *CoreV1Client) *backupPoints {
	return &backupPoints{client: c.RESTClient()}
}

func (c *backupPoints) Get(ctx context.Context, name string, opts v1.GetOptions) (result *corev1.BackupPoint, err error) {
	result = &corev1.BackupPoint{}
	err = c.client.Get().
		Resource("backuppoints").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *backupPoints) List(ctx context.Context, opts v1.ListOptions) (result *corev1.BackupPointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &corev1.BackupPointList{}
	err = c.client.Get().
		Resource("backuppoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *backupPoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("backuppoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
