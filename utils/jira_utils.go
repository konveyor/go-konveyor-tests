package utils

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/onsi/gomega"
)

type Jira api.Tracker

func CreateJiraInstance(data data.JiraInstanceTC) (api.Identity, Jira) {
	jiraIdentity := data.Identity
	uniq.IdentityName(&jiraIdentity)
	err := Identity.Create(&jiraIdentity)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	jiraInstance := Jira{
		URL:  data.JiraUrl,
		Kind: data.JiraKind,
		Identity: api.Ref{
			ID:   jiraIdentity.ID,
			Name: jiraIdentity.Name,
		},
	}
	uniq.TrackerName((*api.Tracker)(&jiraInstance))
	err = Tracker.Create((*api.Tracker)(&jiraInstance))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// Wait for connection succeeded
	var jira *api.Tracker
	var retry, _ = strconv.Atoi(conf.RETRY_NUM)
	for i := 0; i < retry; i++ {
		jira, err = Tracker.Get(jiraInstance.ID)
		if err != nil || jira.Connected {
			break
		}
		time.Sleep(5 * time.Second)
	}
	gomega.Expect(jira.Connected).To(gomega.BeTrue())

	return jiraIdentity, jiraInstance
}

func (r *Jira) DeleteJiraIssues(issues []string) {
	for i := 0; i < len(issues); i++ {
		url := conf.JIRA_CLOUD_URL + "/rest/api/3/issue/" + issues[i]
		r.sendJiraRequest(url, "DELETE")
	}
}

func (r *Jira) sendJiraRequest(url string, method string) {
	// Create a basic authentication string
	basicAuth := "Basic " +
		base64.StdEncoding.EncodeToString([]byte(conf.JIRA_CLOUD_USERNAME+":"+conf.JIRA_CLOUD_PASSWORD))

	// Create a bearer authentication string
	bearerAuth := "Bearer " + conf.JIRA_CLOUD_PASSWORD

	// Create a request
	request, _ := http.NewRequest(method, url, nil)

	// Set the authorization header
	request.Header.Set("Authorization", basicAuth)
	if r.Kind == "on-permise" {
		request.Header.Set("Authorization", bearerAuth)
	}

	// Make the request
	client := http.Client{}
	response, err := client.Do(request)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// Ensure that the response body is closed when the function returns
	defer response.Body.Close()
}
