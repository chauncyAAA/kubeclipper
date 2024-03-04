package main

import (
	"fmt"
	"os"

	"github.com/kubeclipper/kubeclipper/cmd/kubeclipper-agent/app"
	_ "github.com/kubeclipper/kubeclipper/pkg/component/metallb"
	_ "github.com/kubeclipper/kubeclipper/pkg/component/nfs"
	_ "github.com/kubeclipper/kubeclipper/pkg/component/nfscsi"
	_ "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1/cri"
	_ "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1/k8s"
)

func main() {
	cmds := app.NewKiAgentCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := cmds.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
