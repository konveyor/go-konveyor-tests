package go_konveyor_tests_test

import (
	"testing"

	_ "github.com/konveyor/go-konveyor-tests/e2e/centralconfig"
	_ "github.com/konveyor/go-konveyor-tests/e2e/jiraintegration"
	_ "github.com/konveyor/go-konveyor-tests/e2e/metrics"
	_ "github.com/konveyor/go-konveyor-tests/e2e/migrationwave"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoKonveyorTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Konveyor Tests Suite")
}
