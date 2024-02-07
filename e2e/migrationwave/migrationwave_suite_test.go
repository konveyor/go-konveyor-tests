package migrationwave

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMigrationwave(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Migrationwave Suite")
}
