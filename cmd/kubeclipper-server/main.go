package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kubeclipper/kubeclipper/cmd/kubeclipper-server/app"
	_ "github.com/kubeclipper/kubeclipper/pkg/authentication/identityprovider/oidc"
	_ "github.com/kubeclipper/kubeclipper/pkg/clustermanage/kubeadm"
	_ "github.com/kubeclipper/kubeclipper/pkg/component/metallb"
	_ "github.com/kubeclipper/kubeclipper/pkg/component/nfs"
	_ "github.com/kubeclipper/kubeclipper/pkg/component/nfscsi"
	_ "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1/cri"
	_ "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1/k8s"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmds := app.NewKiServerCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := cmds.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
