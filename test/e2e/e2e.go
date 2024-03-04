package e2e

import (
	"testing"

	"github.com/kubeclipper/kubeclipper/test/framework"
	"github.com/kubeclipper/kubeclipper/test/framework/ginkgowrapper"

	"github.com/onsi/ginkgo/v2"
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
	configuration, _ := ginkgo.GinkgoConfiguration()
	framework.Logf("Starting e2e run on Ginkgo %d node", configuration.ParallelProcess)
	ginkgo.RunSpecs(t, "KubeClipper e2e suite")
}
