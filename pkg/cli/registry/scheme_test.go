package registry

import (
	"os"
	"testing"

	"github.com/spf13/cobra"

	"github.com/kubeclipper/kubeclipper/pkg/cli/printer"
)

func TestImage_Printer(t *testing.T) {
	image := &Image{
		Name: "etcd",
		Tags: []string{"v1.1.1", "v2.2.2", "v3.3.3"},
	}
	cmd := &cobra.Command{}
	p := printer.NewPrintFlags()
	p.AddFlags(cmd)
	cmd.Flags().Set("output", "table")
	p.Print(image, os.Stdout)
	cmd.Flags().Set("output", "json")
	p.Print(image, os.Stdout)
	cmd.Flags().Set("output", "yaml")
	p.Print(image, os.Stdout)
}

func TestRepository_Printer(t *testing.T) {
	repos := &Repositories{
		Repositories: []string{
			"caas4/nfs-subdir-external-provisioner",
			"calico/cni",
			"pause"},
	}
	cmd := &cobra.Command{}
	p := printer.NewPrintFlags()
	p.AddFlags(cmd)
	cmd.Flags().Set("output", "table")
	p.Print(repos, os.Stdout)
	cmd.Flags().Set("output", "json")
	p.Print(repos, os.Stdout)
	cmd.Flags().Set("output", "yaml")
	p.Print(repos, os.Stdout)
}
