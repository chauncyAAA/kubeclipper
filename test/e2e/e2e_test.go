package e2e

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/kubeclipper/kubeclipper/test/e2e/cluster"
	_ "github.com/kubeclipper/kubeclipper/test/e2e/node"
	_ "github.com/kubeclipper/kubeclipper/test/e2e/region"
	"github.com/kubeclipper/kubeclipper/test/framework"
)

// handleFlags sets up all flags and parses the command line.
func handleFlags() {
	framework.RegisterCommonFlags(flag.CommandLine)
	flag.Parse()
}

func TestMain(m *testing.M) {
	handleFlags()
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestRunE2ETests(t *testing.T) {
	RunE2ETests(t)
}
