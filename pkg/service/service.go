package service

import (
	"context"
	"time"

	"github.com/kubeclipper/kubeclipper/pkg/oplog"

	"go.uber.org/zap"

	"github.com/kubeclipper/kubeclipper/pkg/logger"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

type Runnable interface {
	PrepareRun(stopCh <-chan struct{}) error
	Run(stopCh <-chan struct{}) error
	Close()
}

type Interface interface {
	Runnable
}

type IDelivery interface {
	DeliverLogRequest(ctx context.Context, operation *LogOperation) (oplog.LogContentResponse, error) // request & response synchronously.
	CmdDelivery
}

type CmdDelivery interface {
	DeliverTaskOperation(ctx context.Context, operation *v1.Operation, opts *Options) error
	DeliverStep(ctx context.Context, operation *v1.Step, opts *Options) error
	DeliverCmd(ctx context.Context, toNode string, cmds []string, timeout time.Duration) ([]byte, error)
}

func HandlerCrash() {
	if r := recover(); r != nil {
		logger.Error("handler crash", zap.Any("recover_for", r))
	}
}
