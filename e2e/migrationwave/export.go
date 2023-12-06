package migrationwave

import (
	"strconv"
	"time"

	"github.com/konveyor/go-konveyor-tests/config"
	"github.com/konveyor/go-konveyor-tests/data"
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

var _ = Describe("Export applications", func() {
	AfterEach(func() {
		// Resources cleanup
		if keep, _ := strconv.ParseBool(config.Config.KEEP); keep {
			return
		}
		Expect(utils.Tracker.Delete(jiraInstance.ID)).To(Succeed())
		Expect(utils.Identity.Delete(jiraIdentity.ID)).To(Succeed())
		Expect(utils.MigrationWave.Delete((migrationWave.ID))).To(Succeed()) /* Associated api.Ticket objects are removed as well */
		jiraInstance.DeleteJiraIssues(issueIds)
		utils.DeleteApplicationsBySlice(appsToExport)
	})

	DescribeTable("",
		func(testCase data.ExportApplicationsCase) {
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
				var retry, _ = strconv.Atoi(config.Config.RETRY_NUM)
				for i := 0; i < retry; i++ {
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
		Entry("Export applications as a task", data.ExportApplicationsCase{
			JiraInstance: data.JiraCloud,
			NumOfApps:    3,
			TicketKind:   "10007", /* Task issuetypeId */
			TicketParent: "10001" /* mta_integration projectId */}))
})
