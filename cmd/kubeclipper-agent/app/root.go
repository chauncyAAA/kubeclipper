package app

import (
	"io"

	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewKiAgentCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:           "kubeclipper-agent",
		Short:         "kubeclipper-agent: k8s installer agent",
		Long:          "TODO: Add long description for kubeclipper-agent",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	stopCh := genericapiserver.SetupSignalHandler()

	cmds.ResetFlags()
	cmds.AddCommand(newCmdVersion(out))
	cmds.AddCommand(newServeCommand(stopCh))

	return cmds
}
