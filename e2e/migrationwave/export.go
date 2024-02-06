package migrationwave

import (
	"os"
	"strconv"
	"time"

	"github.com/konveyor/go-konveyor-tests/data/jira"
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
		func(testCase migrationwave.ExportApplicationsCase) {
			appsToExport = utils.CreateMultipleApplications(testCase.NumOfApps)
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

			By("Check ticket was created")
			for i := 0; i < len(appsToExport); i++ {
				ticket := api.Ticket{Kind: testCase.TicketKind, Parent: testCase.TicketParent, Application: api.Ref{ID: appsToExport[i].ID},
					Tracker: api.Ref{ID: jiraInstance.ID}}
				utils.Ticket.Create(&ticket)

				// Wait for reference field to be populated
				var got *api.Ticket
				for i := 0; i < RETRY_NUM; i++ {
					got, err = utils.Ticket.Get(ticket.ID)
					if err != nil || got.Reference != "" {
						break
					}
					time.Sleep(5 * time.Second)
				}
				Expect(got.Reference).NotTo(BeEmpty())
				issueIds = append(issueIds, got.Reference)
			}

		},
		Entry("Export as task to jira cloud", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraCloud,
			NumOfApps:    3,
			TicketKind:   "10007", /* Task issuetypeId */
			TicketParent: "10001" /* mta_integration projectId */}),
		Entry("Export as story to jira cloud", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraCloud,
			NumOfApps:    2,
			TicketKind:   "10011", /* Story */
			TicketParent: "10001"}),
		Entry("Export as bug to jira cloud", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraCloud,
			NumOfApps:    2,
			TicketKind:   "10012", /* Bug */
			TicketParent: "10001"}),

		Entry("Export as task to jira server", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraServer,
			NumOfApps:    3,
			TicketKind:   "3", /* Task issuetypeId */
			TicketParent: "12340621" /* mta-qe-test projectId */}),
		Entry("Export as story to jira server", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraServer,
			NumOfApps:    2,
			TicketKind:   "17", /* Story */
			TicketParent: "12340621"}),
		Entry("Export as bug to jira server", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraServer,
			NumOfApps:    2,
			TicketKind:   "1", /* Bug */
			TicketParent: "12340621"}),

		Entry("Export as task to jira server using token", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraServerBearerToken, /* Using token for Jira connection */
			NumOfApps:    3,
			TicketKind:   "3", /* Task issuetypeId */
			TicketParent: "12340621" /* mta-qe-test projectId */}),
		Entry("Export as story to jira server using token", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraServerBearerToken,
			NumOfApps:    2,
			TicketKind:   "17", /* Story */
			TicketParent: "12340621"}),
		Entry("Export as bug to jira server using token", migrationwave.ExportApplicationsCase{
			JiraInstance: jira.JiraServerBearerToken, /* Using token for Jira connection */
			NumOfApps:    2,
			TicketKind:   "1", /* Bug */
			TicketParent: "12340621"}))
})
