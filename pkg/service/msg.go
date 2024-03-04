package service

import (
	"time"

	"github.com/kubeclipper/kubeclipper/pkg/errors"
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

type Operation int32

const (
	OperationRegisterNode Operation = iota + 1
	OperationReportNodeStatus
	OperationGetNode
	OperationUpdateNodeLease
	OperationGetNodeLease
	OperationCreateNodeLease
	// task operation
	OperationRunTask
	OperationStepLog
	OperationBackup
	OperationRecovery
	OperationRunCmd
	OperationRunStep
)

const (
	MsgSubjectFormat = "%s.%s"
	// action:bakFileName:opID:stepID
	MsgCreateBackupFormat = "%s:%s:%s:%s"
	// action:bakFileName:id
	MsgDeleteBackupFormat = "%s:%s:%s"
	// downloadDir:filename:id
	MsgStepRecoveryFormat = "%s:%s:%s"
)

type NodeStatusPayload struct {
	Op       Operation `json:"op,omitempty"`
	NodeName string    `json:"node_name,omitempty"`
	Data     []byte    `json:"data,omitempty"`
}

type CommonReply struct {
	Error *errors.StatusError `json:"error,omitempty"`
	Data  []byte              `json:"data,omitempty"`
}

type MsgPayload struct {
	Op                Operation `json:"op,omitempty"`
	OperationIdentity string    `json:"operationIdentity"`
	LastTaskReply     []byte    `json:"lastTaskReply,omitempty"`
	DryRun            bool      `json:"dryRun,omitempty"`
	Retry             bool      `json:"retry,omitempty"`
	Step              v1.Step   `json:"step,omitempty"`
	Cmds              []string  `json:"cmds,omitempty"`
}

type LogOperation struct {
	Op                Operation
	OperationIdentity string // operation ID
	To                string // the node that message will be sent to
	Timeout           time.Duration
}

type Options struct {
	DryRun         bool
	ForceSkipError bool
}
