package migrationwave

import (
	"os"
	"strconv"

	"github.com/konveyor/go-konveyor-tests/data/migrationwave"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/go-konveyor-tests/utils"
	"github.com/konveyor/tackle2-hub/api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	jiraInstance  utils.Jira
	jiraIdentity  api.Identity
	migrationWave api.MigrationWave
	issueIds      []string
	appsToExport  []api.Application
)

const (
	RETRY_NUM = 10
)

var _ = Describe("Export applications", func() {
	var (
		jiraInstance  utils.Jira
		jiraIdentity  api.Identity
		migrationWave api.MigrationWave
		issueIds      []string
		appsToExport  []api.Application
	)

	BeforeEach(func() {
		jiraInstance = utils.Jira{}
		jiraIdentity = api.Identity{}
		migrationWave = api.MigrationWave{}
		issueIds = []string{}
		appsToExport = []api.Application{}

	})
	AfterEach(func() {
		// Resources cleanup
		if keep, _ := strconv.ParseBool(os.Getenv("KEEP")); keep {
			return
		}
		Expect(utils.Tracker.Delete(jiraInstance.ID)).To(Succeed())
		Expect(utils.Identity.Delete(jiraIdentity.ID)).To(Succeed())
		Expect(utils.MigrationWave.Delete((migrationWave.ID))).To(Succeed()) /* Associated api.Ticket objects are removed as well */
		jiraInstance.DeleteJiraIssues(issueIds)
		utils.DeleteApplicationsBySlice(appsToExport)
	})

	DescribeTable("",
		func(testCase migrationwave.ExportApplicationsCase, numOfApps int) {
			appsToExport = utils.CreateMultipleApplications(numOfApps)
			By("Create migration wave")
			uniq.MigrationWaveName(&migrationWave)
			err := utils.MigrationWave.Create(&migrationWave)
			Expect(err).NotTo(HaveOccurred())

			By("Manage applications")
			for i := 0; i < len(appsToExport); i++ {
				migrationWave.Applications = append(migrationWave.Applications, api.Ref{ID: appsToExport[i].ID, Name: appsToExport[i].Name})
			}
			err = utils.MigrationWave.Update(&migrationWave)
			Expect(err).NotTo(HaveOccurred())

			By("Create Jira instance")
			jiraIdentity, jiraInstance = utils.CreateJiraInstance(testCase.JiraInstance)
			jiraInstance.CheckConnection()

			By("Create ticket, check ticket was created")
			for i := 0; i < len(appsToExport); i++ {
				ticket := api.Ticket{Kind: testCase.TicketKind, Parent: testCase.TicketParent, Application: api.Ref{ID: appsToExport[i].ID},
					Tracker: api.Ref{ID: jiraInstance.ID}}
				utils.Ticket.Create(&ticket)

				ticketReference := utils.CheckReferenceNotEmpty(&ticket)
				issueIds = append(issueIds, ticketReference)
			}
		},
		Entry("Export as task to jira cloud", migrationwave.ExportBugToJiraCloud, 3),
		Entry("Export as story to jira cloud", migrationwave.ExportTaskToJiraCloud, 2),
		Entry("Export as bug to jira cloud", migrationwave.ExportBugToJiraCloud, 2),

		Entry("Export as task to jira server", migrationwave.ExportTaskToJiraServer, 3),
		Entry("Export as story to jira server", migrationwave.ExportStoryToJiraServer, 2),
		Entry("Export as bug to jira server", migrationwave.ExportBugToJiraServer, 2),

		Entry("Export as task to jira server using token", migrationwave.ExportTaskToJiraUsingToken, 3),
		Entry("Export as story to jira server using token", migrationwave.ExportStoryToJiraUsingToken, 2),
		Entry("Export as bug to jira server using token", migrationwave.ExportBugToJiraUsingToken, 2))
})
