package e2e

import (
	"testing"

	"github.com/kubeclipper/kubeclipper/test/framework"
	"github.com/kubeclipper/kubeclipper/test/framework/ginkgowrapper"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/gomega"
)

var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Reference common test to make the import valid.
	//setupSuite()
	return nil
}, func(data []byte) {
	// Run on all Ginkgo nodes
	//setupSuitePerGinkgoNode()
})

var _ = ginkgo.SynchronizedAfterSuite(func() {
	CleanupSuite()
}, func() {
	AfterSuiteActions()
})

func RunE2ETests(t *testing.T) {
	gomega.RegisterFailHandler(ginkgowrapper.Fail)
	framework.Logf("Starting e2e run on Ginkgo %d node", config.GinkgoConfig.ParallelNode)
	ginkgo.RunSpecs(t, "KubeClipper e2e suite")
}
