//go:generate mockgen -destination mock/mock_nats.go -source interface.go Interface

package natsio

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type Msg struct {
	Subject string
	From    string
	To      string
	Step    string
	Timeout time.Duration
	Data    []byte
}

type ReplyHandler func(msg *nats.Msg) error
type TimeoutHandler func(msg *Msg) error

type Interface interface {
	SetDisconnectErrHandler(handler nats.ConnErrHandler)
	SetReconnectHandler(handler nats.ConnHandler)
	SetErrorHandler(handler nats.ErrHandler)
	SetClosedHandler(handler nats.ConnHandler)
	RunServer(stopCh <-chan struct{}) error
	InitConn(stopCh <-chan struct{}) error
	Publish(msg *Msg) error
	Subscribe(subj string, handler nats.MsgHandler) error
	QueueSubscribe(subj string, queue string, handler nats.MsgHandler) error
	Request(msg *Msg, timeoutHandler TimeoutHandler) ([]byte, error)
	RequestWithContext(ctx context.Context, msg *Msg) ([]byte, error)
	RequestAsync(msg *Msg, handler ReplyHandler, timeoutHandler TimeoutHandler) error
	Close()
}
