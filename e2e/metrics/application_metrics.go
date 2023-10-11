package metrics

import (
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/konveyor/go-konveyor-tests/config"
	"github.com/konveyor/go-konveyor-tests/utils"
	"github.com/konveyor/tackle2-hub/api"
)

var (
	metricValue int
	metricName  string = "konveyor_applications_inventoried"
	apps        []api.Application
	Application = utils.Application
)

var _ = Describe("Application Metrics", Ordered, func() {
	BeforeAll(func() {
		// Get the current metric value
		metricValue = GetMetricValue(metricName)
	})

	AfterAll(func() {
		// Resources cleanup
		if keep, _ := strconv.ParseBool(Config.KEEP); keep {
			return
		}
		utils.DeleteApplicationsBySlice(apps)
	})

	Context("Create applications", func() {
		It("Should increase the applications gauge", func() {
			numOfApps := 3
			apps = utils.CreateMultipleApplications(numOfApps)
			metricValue += len(apps)
			Expect(GetMetricValue(metricName)).To(Equal(metricValue))
		})
	})

	Context("Delete application", func() {
		It("Should decrease the applications gauge", func() {
			// Delete the last application item
			lastApplicationItem := apps[len(apps)-1]
			utils.DeleteApplicationsBySlice([]api.Application{lastApplicationItem})
			metricValue--
			Expect(GetMetricValue(metricName)).To(Equal(metricValue))
		})
	})

})
