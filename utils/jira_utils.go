package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/konveyor/go-konveyor-tests/data/jira"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/onsi/gomega"
)

type Jira api.Tracker

func RandomJiraData() (jiraData jira.JiraInstanceTC) {
	randomIndex := rand.Intn(len(jira.JiraSamples))
	//// Create a local copy
	jiraData = jira.JiraSamples[randomIndex]
	return
}

func CreateJiraInstance(data jira.JiraInstanceTC) (api.Identity, Jira) {
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

	return jiraIdentity, jiraInstance
}

func (r *Jira) CheckConnection() {
	// Wait for connection succeeded
	var jira *api.Tracker
	for i := 0; i < RETRY_NUM; i++ {
		jira, err := Tracker.Get(r.ID)
		if err != nil || jira.Connected {
			break
		}
		time.Sleep(5 * time.Second)
	}
	gomega.Expect(jira.Connected).To(gomega.BeTrue())
}

func (r *Jira) DeleteJiraIssues(issues []string) {
	if r.Kind == jira.JIRA_KIND_CLOUD {
		for i := 0; i < len(issues); i++ {
			url := r.URL + "/rest/api/3/issue/" + issues[i]
			r.sendJiraRequest(url, "DELETE")
		}
		return
	}
	if r.Kind == jira.JIRA_KIND_ONPREM {
		for i := 0; i < len(issues); i++ {
			url := fmt.Sprintf("%s/rest/api/2/issue/%s/archive", r.URL, issues[i])
			r.sendJiraRequest(url, "PUT")
		}
		return
	}
}

func (r *Jira) sendJiraRequest(url string, method string) {
	// Create a basic authentication string
	basicAuth := "Basic " +
		base64.StdEncoding.EncodeToString([]byte(os.Getenv("JIRA_CLOUD_USERNAME")+":"+os.Getenv("JIRA_CLOUD_PASSWORD")))

	// Create a bearer authentication string
	bearerAuth := "Bearer " + os.Getenv("JIRA_SERVER_TOKEN")

	// Create a request
	request, _ := http.NewRequest(method, url, nil)

	// Set the authorization header
	request.Header.Set("Authorization", basicAuth)
	if r.Kind == jira.JIRA_KIND_ONPREM {
		request.Header.Set("Authorization", bearerAuth)
	}

	// Make the request
	client := http.Client{}
	response, err := client.Do(request)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	// Ensure that the response body is closed when the function returns
	defer response.Body.Close()
}
