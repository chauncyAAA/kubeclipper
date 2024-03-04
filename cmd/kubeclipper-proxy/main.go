package main

import (
	"fmt"
	"os"

	"github.com/kubeclipper/kubeclipper/cmd/kubeclipper-proxy/app"
	_ "github.com/kubeclipper/kubeclipper/pkg/component/nfs"
	_ "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1/cri"
	_ "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1/k8s"
)

func main() {
	cmds := app.NewKiProxyCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := cmds.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
