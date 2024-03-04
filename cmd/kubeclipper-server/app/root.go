package app

import (
	"io"

	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewKiServerCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:           "kubeclipper-server",
		Short:         "kubeclipper-server: k8s installer server",
		Long:          "TODO: Add long description for kubeclipper-server",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	stopCh := genericapiserver.SetupSignalHandler()

	cmds.ResetFlags()
	cmds.CompletionOptions.DisableDefaultCmd = true
	cmds.AddCommand(newCmdVersion(out))
	cmds.AddCommand(newServeCommand(stopCh))

	return cmds
}
