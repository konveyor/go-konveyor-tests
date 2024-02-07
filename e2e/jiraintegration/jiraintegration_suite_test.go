package jiraintegration

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJiraintegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jiraintegration Suite")
}
