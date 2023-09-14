package jiraintegration

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"strconv"
	"time"

	. "github.com/konveyor/go-konveyor-tests/config"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	. "github.com/konveyor/go-konveyor-tests/utils"
	"github.com/konveyor/tackle2-hub/api"
)

var (
	jiraBasicAuth api.Identity
	jiraInstance  api.Tracker
)

const (
	jiraBasic string = "basic-auth"
	jiraCloud string = "jira-cloud"
)

var _ = Describe("Jira Connection", Ordered, func() {
	AfterAll(func() {
		// Resources cleanup
		// Delete tracker instance and the associated identity after
		if keep, _ := strconv.ParseBool(Config.KEEP); keep {
			return
		}
		Expect(Tracker.Delete(jiraInstance.ID)).To(Succeed())
		Expect(Identity.Delete(jiraBasicAuth.ID)).To(Succeed())
	})

	Context("Create new Jira instance", func() {
		It("Verify connection", func() {
			By("Create credential")
			jiraBasicAuth = api.Identity{
				Kind:     jiraBasic,
				User:     Config.JIRA_USERNAME,
				Password: Config.JIRA_PASSWORD,
			}
			uniq.IdentityName(&jiraBasicAuth)
			err := Identity.Create(&jiraBasicAuth)
			Expect(err).NotTo(HaveOccurred())

			By("Create Jira cloud instance and associate credential")
			jiraInstance = api.Tracker{
				URL:  Config.JIRA_URL,
				Kind: jiraCloud,
				Identity: api.Ref{
					ID:   jiraBasicAuth.ID,
					Name: jiraBasicAuth.Name,
				},
			}
			uniq.TrackerName(&jiraInstance)
			err = Tracker.Create(&jiraInstance)
			Expect(err).NotTo(HaveOccurred())

			// Wait for connection succeeded
			var jira *api.Tracker
			var retry, _ = strconv.Atoi(Config.RETRY_NUM)
			for i := 0; i < retry; i++ {
				jira, err = Tracker.Get(jiraInstance.ID)
				if err != nil || jira.Connected {
					break
				}
				time.Sleep(5 * time.Second)
			}
			Expect(jira.Connected).To(BeTrue())
		})
	})

})
