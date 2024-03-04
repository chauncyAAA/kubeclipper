package framework

import (
	"flag"

	"github.com/kubeclipper/kubeclipper/pkg/constatns"
	"github.com/onsi/ginkgo/v2"
)

const (
	defaultHost          = "http://127.0.0.1:8080"
	defaultServiceSubnet = constatns.ClusterServiceSubnet
	defaultPodSubnet     = constatns.ClusterPodSubnet
	defaultLocalRegistry = "127.0.0.1:5000"
	defaultWorkerNodeVip = "169.254.169.100"
)

type TestContextType struct {
	Host          string
	InMemoryTest  bool
	ServiceSubnet string
	PodSubnet     string
	LocalRegistry string
	WorkerNodeVip string
}

// TestContext should be used by all tests to access common context data.
var TestContext TestContextType

func RegisterCommonFlags(flags *flag.FlagSet) {
	suiteConfig, reporterConfig := ginkgo.GinkgoConfiguration()
	// Turn on verbose by default to get spec names
	reporterConfig.Verbose = true

	// Randomize specs as well as suites
	suiteConfig.RandomizeAllSpecs = true

	flag.BoolVar(&TestContext.InMemoryTest, "in-memory-test", false,
		"Whether Ki-server and Ki-agent be started in memory.")
	flag.StringVar(&TestContext.Host, "server-address", defaultHost,
		"Ki Server API Server IP/DNS, default 127.0.0.1:8080")
	flag.StringVar(&TestContext.ServiceSubnet, "svc-subnet", defaultServiceSubnet,
		"cluster svc sub net, default 10.96.0.0/12")
	flag.StringVar(&TestContext.PodSubnet, "pod-subnet", defaultPodSubnet,
		"cluster pod sub net, default 172.25.0.0/16")
	flag.StringVar(&TestContext.LocalRegistry, "registry", defaultLocalRegistry,
		"cri image registry addr, default 127.0.0.1:5000")
	flag.StringVar(&TestContext.WorkerNodeVip, "vip", defaultWorkerNodeVip,
		"cluster worker node loadblance vip, default 169.254.169.100")
	flag.DurationVar(&clusterInstallShort, "cluster-install-short-timeout", clusterInstallShort,
		"cluster install short timeout interval")
}
