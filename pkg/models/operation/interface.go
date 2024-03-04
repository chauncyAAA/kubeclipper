//go:generate mockgen -destination mock/mock_operation.go -source interface.go Operator

package operation

import (
	"context"

	"github.com/kubeclipper/kubeclipper/pkg/models"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/kubeclipper/kubeclipper/pkg/query"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

type Reader interface {
	ListOperations(ctx context.Context, query *query.Query) (*v1.OperationList, error)
	WatchOperations(ctx context.Context, query *query.Query) (watch.Interface, error)
	GetOperation(ctx context.Context, name string) (*v1.Operation, error)
	ReaderEx
}

type ReaderEx interface {
	GetOperationEx(ctx context.Context, name string, resourceVersion string) (*v1.Operation, error)
	ListOperationsEx(ctx context.Context, query *query.Query) (*models.PageableResponse, error)
}

type Writer interface {
	DeleteOperation(ctx context.Context, name string) error
	CreateOperation(ctx context.Context, operation *v1.Operation) (*v1.Operation, error)
	UpdateOperation(ctx context.Context, operation *v1.Operation) (*v1.Operation, error)
	UpdateOperationStatus(ctx context.Context, name string, status *v1.OperationStatus) (*v1.Operation, error)
	DeleteOperationCollection(ctx context.Context, query *query.Query) error
}

type Operator interface {
	Reader
	Writer
}
