package metrics

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/konveyor/go-konveyor-tests/application"
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

	AfterEach(func() {
		// Cleanup resources
	})

	Context("when creating application", func() {
		It("should increase the applications gauge", func() {
			numOfApps := 3
			apps := CreateMultipleApplications(numOfApps)
			expectedValue := metricValue + len(apps)
			Expect(GetMetricValue(metricName)).To(Equal(expectedValue))
		})
	})

})
