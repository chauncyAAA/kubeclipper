package kcctl

import "github.com/onsi/ginkgo"

// SIGDescribe annotates the test with the SIG label.
func SIGDescribe(text string, body func()) bool {
	return ginkgo.Describe("[sig-kcctl] "+text, body)
}
