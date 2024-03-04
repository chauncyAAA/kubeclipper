package component

import (
	"errors"
	"strings"
)

var (
	_agentSteps = defaultAgentStepHandler()
)

var (
	ErrStepExist     = errors.New("component already exist")
	ErrStepKeyFormat = errors.New("component key must be name/version")
)

const OfflinePackagesKeyFormat = "%s-%s-%s"

type agentStep struct {
	steps map[string]StepRunnable
}

func defaultAgentStepHandler() agentStep {
	return agentStep{steps: map[string]StepRunnable{}}
}

// RegisterAgentStep KV must format at componentName/version/stepName
func RegisterAgentStep(kv string, p StepRunnable) error {
	if !checkAgentStepKey(kv) {
		return ErrStepKeyFormat
	}
	return _agentSteps.registerAgentStep(kv, p)
}

func LoadAgentStep(kv string) (StepRunnable, bool) {
	return _agentSteps.load(kv)
}

func (h *agentStep) load(kv string) (StepRunnable, bool) {
	c, exist := h.steps[kv]
	return c, exist
}

func (h *agentStep) registerAgentStep(kv string, p StepRunnable) error {
	_, exist := h.steps[kv]
	if exist {
		return ErrStepExist
	}
	h.steps[kv] = p
	return nil
}

func checkAgentStepKey(kv string) bool {
	parts := strings.Split(kv, "/")
	return len(parts) == 3
}
