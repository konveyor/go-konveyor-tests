package jira

import (
	"os"

	"github.com/konveyor/tackle2-hub/api"
)

type JiraInstanceTC struct {
	Identity api.Identity
	JiraUrl  string
	JiraKind string
}

const (
	JIRA_KIND_CLOUD  = "jira-cloud"
	JIRA_KIND_ONPREM = "jira-onprem"
)

// Set of valid instances for tests and reuse.
// Important: Do not use it directly to not affect other tests.
var (
	JiraCloud = JiraInstanceTC{
		Identity: api.Identity{
			Kind:     "basic-auth",
			User:     os.Getenv("JIRA_CLOUD_USERNAME"),
			Password: os.Getenv("JIRA_CLOUD_PASSWORD"),
		},
		JiraUrl:  os.Getenv("JIRA_CLOUD_URL"),
		JiraKind: JIRA_KIND_CLOUD,
	}
	JiraServer = JiraInstanceTC{
		Identity: api.Identity{
			Kind:     "basic-auth",
			User:     os.Getenv("JIRA_SERVER_USERNAME"),
			Password: os.Getenv("JIRA_SERVER_PASSWORD"),
		},
		JiraUrl:  os.Getenv("JIRA_SERVER_URL"),
		JiraKind: JIRA_KIND_ONPREM,
	}
	JiraServerBearerToken = JiraInstanceTC{
		Identity: api.Identity{
			Kind: "bearer",
			Key:  os.Getenv("JIRA_SERVER_TOKEN"),
		},
		JiraUrl:  os.Getenv("JIRA_SERVER_URL"),
		JiraKind: JIRA_KIND_ONPREM,
	}

	JiraSamples = []JiraInstanceTC{JiraCloud, JiraServer, JiraServerBearerToken}
)
