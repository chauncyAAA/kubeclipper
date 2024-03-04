package app

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/component-base/version"
	"sigs.k8s.io/yaml"
)

// newCmdVersion provides the version information of kubeclipper-server.
func newCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of kubeclipper-server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunVersion(out, cmd)
		},
		Args: cobra.NoArgs,
	}
	cmd.Flags().StringP("output", "o", "", "Output format; available options are 'yaml', 'json' and 'short'")
	return cmd
}

// RunVersion provides the version information of kubeclipper-server in format depending on arguments
// specified in cobra.Command.
func RunVersion(out io.Writer, cmd *cobra.Command) error {
	v := version.Get()

	const flag = "output"
	of, err := cmd.Flags().GetString(flag)
	if err != nil {
		return errors.Wrapf(err, "error accessing flag %s for command %s", flag, cmd.Name())
	}

	switch of {
	case "":
		fmt.Fprintf(out, "kubeclipper server version: %#v\n", v)
	case "short":
		fmt.Fprintf(out, "%s\n", v.GitVersion)
	case "yaml":
		y, err := yaml.Marshal(&v)
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(y))
	case "json":
		y, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(y))
	default:
		return errors.Errorf("invalid output format: %s", of)
	}

	return nil
}
