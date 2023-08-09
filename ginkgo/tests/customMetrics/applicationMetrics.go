package customMetrics

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/konveyor/go-konveyor-tests/ginkgo/utils"
	. "github.com/konveyor/go-konveyor-tests/ginkgo/utils/migration/applicationInventory"
)

var (
	metricValue int
	metricName  string = "konveyor_applications_inventoried"
)

var _ = Describe("Application Metrics", func() {
	BeforeEach(func() {
		// Get the current metric value
		metricValue = GetMetricValue(metricName)
	})

	Context("when creating application", func() {
		It("should increase the applications gauge", func() {
			apps := CreateMultipleApplications(3)
			expectedValue := metricValue + len(apps)
			Expect(GetMetricValue(metricName)).To(Equal(expectedValue))
		})
	})

})
