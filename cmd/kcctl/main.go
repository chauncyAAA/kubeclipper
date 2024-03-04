package main

import (
	"fmt"
	"os"

	"github.com/kubeclipper/kubeclipper/cmd/kcctl/app"
)

func main() {
	cmds := app.NewKubeClipperCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := cmds.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
