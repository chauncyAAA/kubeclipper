package region

import "github.com/onsi/ginkgo/v2"

// SIGDescribe annotates the test with the SIG label.
func SIGDescribe(text string, body func()) bool {
	return ginkgo.Describe("[sig-region] "+text, body)
}
