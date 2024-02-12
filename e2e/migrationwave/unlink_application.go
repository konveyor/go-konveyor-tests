package migrationwave

import (
	"math/rand"
	"os"
	"strconv"

	"github.com/konveyor/go-konveyor-tests/data/migrationwave"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/go-konveyor-tests/utils"
	"github.com/konveyor/tackle2-hub/api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unlink application from Jira", Ordered, func() {
	const (
		numOfApps = 1
	)

	var (
		jiraInstance  utils.Jira
		jiraIdentity  api.Identity
		migrationWave api.MigrationWave
		issueIds      []string
		appsToExport  []api.Application
		tickets       []api.Ticket
	)

	BeforeAll(func() {
		// Random TC
		samples := migrationwave.Samples
		index := rand.Intn(len(samples))
		randomExportTC := samples[index]

		appsToExport = utils.CreateMultipleApplications(numOfApps)
		// Create migration wave
		uniq.MigrationWaveName(&migrationWave)
		utils.MigrationWave.Create(&migrationWave)

		//  Attach applications
		for i := 0; i < len(appsToExport); i++ {
			migrationWave.Applications = append(migrationWave.Applications, api.Ref{ID: appsToExport[i].ID, Name: appsToExport[i].Name})
		}
		utils.MigrationWave.Update(&migrationWave)

		// Create Jira instance
		jiraIdentity, jiraInstance = utils.CreateJiraInstance(randomExportTC.JiraInstance)

		// Create ticket
		for i := 0; i < len(appsToExport); i++ {
			ticket := api.Ticket{Kind: randomExportTC.TicketKind, Parent: randomExportTC.TicketParent,
				Application: api.Ref{ID: appsToExport[i].ID}, Tracker: api.Ref{ID: jiraInstance.ID}}
			utils.Ticket.Create(&ticket)

			ticketReference := utils.WaitForReference(&ticket)
			issueIds = append(issueIds, ticketReference)
			tickets = append(tickets, ticket)
		}
	})

	AfterAll(func() {
		// Resources cleanup
		if keep, _ := strconv.ParseBool(os.Getenv("KEEP")); keep {
			return
		}
		Expect(utils.Identity.Delete(jiraIdentity.ID)).To(Succeed())
		Expect(utils.MigrationWave.Delete((migrationWave.ID))).To(Succeed()) /* Associated api.Ticket objects are removed as well */
		jiraInstance.DeleteJiraIssues(issueIds)
		utils.DeleteApplicationsBySlice(appsToExport)
	})

	Context("Delete jira instance", func() {
		// TODO (mguetta1): Remove description head once bug is fixed
		It("Bug MTA-2226 | Delete fails because an application is linked", func() {
			err := utils.Tracker.Delete(jiraInstance.ID)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Delete tickets", func() {
		It("Ticket should be deleted", func() {
			for _, t := range tickets {
				err := utils.Ticket.Delete(t.ID)
				Expect(err).ShouldNot(HaveOccurred())
			}
		})

		It("Delete jira instance - instance should be deleted", func() {
			err := utils.Tracker.Delete(jiraInstance.ID)
			Expect(err).ShouldNot(HaveOccurred())
		})

	})
})
