package jiraintegration

import (
	"strconv"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/konveyor/go-konveyor-tests/config"
	. "github.com/konveyor/go-konveyor-tests/utils"

	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/tackle2-hub/api"
)

var _ = Describe("Jira connection", func() {
	var jiraIdentity api.Identity
	var jiraInstance api.Tracker

	AfterEach(func() {
		// Resources cleanup
		// Delete tracker instance and the associated identity after
		if keep, _ := strconv.ParseBool(Config.KEEP); keep {
			return
		}
		Expect(Tracker.Delete(jiraInstance.ID)).To(Succeed())
		Expect(Identity.Delete(jiraIdentity.ID)).To(Succeed())
	})

	DescribeTable("",
		func(testCase data.JiraInstanceTC) {
			By("Create credential")
			jiraIdentity = testCase.Identity
			uniq.IdentityName(&jiraIdentity)
			err := Identity.Create(&jiraIdentity)
			Expect(err).NotTo(HaveOccurred())

			By("Create Jira instance and associate credential")
			jiraInstance = api.Tracker{
				URL:  testCase.JiraUrl,
				Kind: testCase.JiraKind,
				Identity: api.Ref{
					ID:   jiraIdentity.ID,
					Name: jiraIdentity.Name,
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

		},
		Entry("Jira cloud", data.JiraCloud),
		Entry("Jira server with basic auth", data.JiraServer),
		Entry("Jira server with bearer token", data.JiraServerBearerToken))
})
