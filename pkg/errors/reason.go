package errors

type StatusReason string

const (
	StatusReasonUnknown           StatusReason = ""
	StatusReasonStorageMethodCall StatusReason = "call etcd storage method error"
	StatusReasonUnexpected        StatusReason = "unexpected error"
)

type CauseType string

const (
	StorageMethodCall  CauseType = "call etcd storage method error"
	Marshal            CauseType = "marshal error"
	Unmarshal          CauseType = "unmarshal error"
	AgentStepInstall   CauseType = "agent install step command "
	AgentStepUninstall CauseType = "agent uninstall step command"
	ShellCommand       CauseType = "shell command step error"
	StepLog            CauseType = "step log error"
)
