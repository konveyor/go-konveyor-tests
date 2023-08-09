package customMetrics_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCustomMetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CustomMetrics Suite")
}
