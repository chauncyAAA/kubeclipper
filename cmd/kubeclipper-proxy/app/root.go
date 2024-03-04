package app

import (
	"io"

	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewKiProxyCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:           "kubeclipper-proxy",
		Short:         "kubeclipper-proxy: kubeclipper proxy server",
		Long:          "TODO: Add long description for kubeclipper-proxy",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	stopCh := genericapiserver.SetupSignalHandler()

	cmds.ResetFlags()
	cmds.AddCommand(newCmdVersion(out))
	cmds.AddCommand(newServeCommand(stopCh))

	return cmds
}
