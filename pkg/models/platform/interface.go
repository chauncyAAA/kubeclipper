//go:generate mockgen -destination mock/mock_platform.go -source interface.go Operator

package platform

import (
	"context"

	"github.com/kubeclipper/kubeclipper/pkg/models"
	"github.com/kubeclipper/kubeclipper/pkg/query"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

type Operator interface {
	Reader
	Writer

	EventReader
	EventWriter
}

type Reader interface {
	GetPlatformSetting(ctx context.Context) (*v1.PlatformSetting, error)
	ReaderEx
}

type Writer interface {
	CreatePlatformSetting(ctx context.Context, platformSetting *v1.PlatformSetting) (*v1.PlatformSetting, error)
	UpdatePlatformSetting(ctx context.Context, platformSetting *v1.PlatformSetting) (*v1.PlatformSetting, error)
}

type ReaderEx interface{}

type EventReader interface {
	ListEvents(ctx context.Context, query *query.Query) (*v1.EventList, error)
	GetEvent(ctx context.Context, name string) (*v1.Event, error)
	EventReaderEx
}

type EventReaderEx interface {
	GetEventEx(ctx context.Context, name string, resourceVersion string) (*v1.Event, error)
	ListEventsEx(ctx context.Context, query *query.Query) (*models.PageableResponse, error)
}

type EventWriter interface {
	CreateEvent(ctx context.Context, Event *v1.Event) (*v1.Event, error)
	DeleteEvent(ctx context.Context, name string) error
	DeleteEventCollection(ctx context.Context, query *query.Query) error
}
