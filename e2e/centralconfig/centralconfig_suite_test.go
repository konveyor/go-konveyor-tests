package centralconfig

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCentralConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Centralized Configuration API Test Suite")
}

var _ = BeforeSuite(func() {
	// Suite setup if needed
})

var _ = AfterSuite(func() {
	// Suite cleanup if needed
})
