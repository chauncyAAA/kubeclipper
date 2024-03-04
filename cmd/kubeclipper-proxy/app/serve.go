package app

import (
	"github.com/spf13/cobra"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/kubeclipper/kubeclipper/pkg/proxy"

	"github.com/kubeclipper/kubeclipper/cmd/kubeclipper-proxy/app/options"
)

func newServeCommand(stopCh <-chan struct{}) *cobra.Command {
	s := options.NewProxyOptions()
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Launch a kubeclipper-proxy",
		Long:  "TODO: add long description for kubeclipper-agent",
		RunE: func(c *cobra.Command, args []string) error {
			err := s.CompletionOptions()
			if err != nil {
				return err
			}
			if errs := s.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}
			return Run(s, stopCh)
		},
		SilenceUsage: true,
	}

	return cmd
}

func Run(s *options.ProxyOptions, stopCh <-chan struct{}) error {
	server, err := proxy.NewServer(s, stopCh)
	if err != nil {
		return err
	}
	err = server.PrepareRun()
	if err != nil {
		return err
	}

	return server.Run()
}
