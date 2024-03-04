package options

import (
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/kubeclipper/kubeclipper/pkg/agent"
	agentconfig "github.com/kubeclipper/kubeclipper/pkg/agent/config"
	"github.com/kubeclipper/kubeclipper/pkg/simple/generic"
)

type AgentOptions struct {
	GenericServerRunOptions *generic.ServerRunOptions
	*agentconfig.Config
}

func NewAgentOptions() *AgentOptions {
	return &AgentOptions{
		GenericServerRunOptions: generic.NewServerRunOptions(),
		Config:                  agentconfig.New(),
	}
}

func (s *AgentOptions) Validate() []error {
	var errors []error
	errors = append(errors, s.GenericServerRunOptions.Validate()...)
	errors = append(errors, s.LogOptions.Validate()...)
	errors = append(errors, s.OpLogOptions.Validate()...)
	errors = append(errors, s.ImageProxyOptions.Validate()...)
	return errors
}

func (s *AgentOptions) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("generic")
	s.GenericServerRunOptions.AddFlags(fs, s.GenericServerRunOptions)
	s.LogOptions.AddFlags(fss.FlagSet("log"))
	s.MQOptions.AddFlags(fss.FlagSet("mq"))
	s.OpLogOptions.AddFlags(fss.FlagSet("oplog"))
	s.ImageProxyOptions.AddFlags(fss.FlagSet("imageProxy"))

	return fss
}

func (s *AgentOptions) NewServer(stopCh <-chan struct{}) (*agent.Server, error) {
	server := &agent.Server{
		Config: s.Config,
	}
	return server, nil
}
