//go:generate mockgen -destination mock/mock_lease.go -source interface.go Operator

package lease

import (
	"context"

	"k8s.io/apimachinery/pkg/watch"

	coordinationv1 "k8s.io/api/coordination/v1"

	"github.com/kubeclipper/kubeclipper/pkg/query"
)

type Operator interface {
	LeaseReader
	LeaseWriter
}

type LeaseReader interface {
	ListLeases(ctx context.Context, query *query.Query) (*coordinationv1.LeaseList, error)
	WatchLease(ctx context.Context, query *query.Query) (watch.Interface, error)
	GetLease(ctx context.Context, name string, resourceVersion string) (*coordinationv1.Lease, error)
	GetLeaseWithNamespace(ctx context.Context, name string, namespace string) (*coordinationv1.Lease, error)
	LeaseReaderEx
}

type LeaseReaderEx interface {
	GetLeaseWithNamespaceEx(ctx context.Context, name string, namespace, resourceVersion string) (*coordinationv1.Lease, error)
	// ListLeasesEx(ctx context.Context, query *query.Query) (*models.PageableResponse, error)
}

type LeaseWriter interface {
	CreateLease(ctx context.Context, lease *coordinationv1.Lease) (*coordinationv1.Lease, error)
	UpdateLease(ctx context.Context, lease *coordinationv1.Lease) (*coordinationv1.Lease, error)
}
