package jiraintegration

import (
	"os"
	"strconv"

	"github.com/konveyor/go-konveyor-tests/data/jira"
	"github.com/konveyor/go-konveyor-tests/utils"
	"github.com/konveyor/tackle2-hub/api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jira connection", func() {
	var jiraIdentity api.Identity
	var jiraInstance utils.Jira

	AfterEach(func() {
		// Resources cleanup
		// Delete tracker instance and the associated identity after
		if keep, _ := strconv.ParseBool(os.Getenv("KEEP")); keep {
			return
		}
		Expect(utils.Tracker.Delete(jiraInstance.ID)).To(Succeed())
		Expect(utils.Identity.Delete(jiraIdentity.ID)).To(Succeed())
	})

	DescribeTable("",
		func(testCase jira.JiraInstanceTC) {
			// Create Jira instance
			jiraIdentity, jiraInstance = utils.CreateJiraInstance(testCase)
		},
		Entry("Jira cloud", jira.JiraCloud),
		Entry("Jira server with basic auth", jira.JiraServer),
		Entry("Jira server with bearer token", jira.JiraServerBearerToken))
})
