package metrics

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/konveyor/go-konveyor-tests/application"
	"github.com/konveyor/tackle2-hub/api"
)

var (
	metricValue int
	metricName  string = "konveyor_applications_inventoried"
	apps        []api.Application
	Application = application.Application
)

var _ = Describe("Application Metrics", Ordered, func() {
	BeforeAll(func() {
		// Get the current metric value
		metricValue = GetMetricValue(metricName)
	})

	AfterAll(func() {
		// Resources cleanup
		application.DeleteBySlice(apps)
	})

	Context("Create applications", func() {
		It("Should increase the applications gauge", func() {
			numOfApps := 3
			apps = application.CreateMultipleApplications(numOfApps)
			metricValue += len(apps)
			Expect(GetMetricValue(metricName)).To(Equal(metricValue))
		})
	})

	Context("Delete application", func() {
		It("Should decrease the applications gauge", func() {
			// Delete the last application item
			lastApplicationItem := apps[len(apps)-1]
			Application.Delete(lastApplicationItem.ID)
			metricValue--
			Expect(GetMetricValue(metricName)).To(Equal(metricValue))
		})
	})

})
