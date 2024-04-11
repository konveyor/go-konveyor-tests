package migrationwave

import (
	"github.com/konveyor/go-konveyor-tests/data/jira"
)

type ExportApplicationsCase struct {
	JiraInstance jira.JiraInstanceTC
	TicketKind   string
	TicketParent string
}

var (
	TaskJiraCloud = ExportApplicationsCase{
		JiraInstance: jira.JiraCloud,
		TicketKind:   "10007", /* Task issuetypeId */
		TicketParent: "10001"} /* mta_integration projectId */

	StoryJiraCloud = ExportApplicationsCase{
		JiraInstance: jira.JiraCloud,
		TicketKind:   "10011", /* Story */
		TicketParent: "10001"}

	BugJiraCloud = ExportApplicationsCase{
		JiraInstance: jira.JiraCloud,
		TicketKind:   "10012", /* Bug */
		TicketParent: "10001"}

	TaskJiraServer = ExportApplicationsCase{
		JiraInstance: jira.JiraServer,
		TicketKind:   "3",        /* Task issuetypeId */
		TicketParent: "12335626"} /* Migration Toolkit for Applications projectId */

	StoryJiraServer = ExportApplicationsCase{
		JiraInstance: jira.JiraServer,
		TicketKind:   "17", /* Story */
		TicketParent: "12335626"}

	BugJiraServer = ExportApplicationsCase{
		JiraInstance: jira.JiraServer,
		TicketKind:   "1", /* Bug */
		TicketParent: "12335626"}

	TaskJiraUsingToken = ExportApplicationsCase{
		JiraInstance: jira.JiraServerBearerToken, /* Using token for Jira connection */
		TicketKind:   "3",                        /* Task issuetypeId */
		TicketParent: "12335626"}

	StoryJiraUsingToken = ExportApplicationsCase{
		JiraInstance: jira.JiraServerBearerToken,
		TicketKind:   "17", /* Story */
		TicketParent: "12335626"}

	BugJiraUsingToken = ExportApplicationsCase{
		JiraInstance: jira.JiraServerBearerToken,
		TicketKind:   "1", /* Bug */
		TicketParent: "12335626"}

	Samples = []ExportApplicationsCase{TaskJiraCloud, TaskJiraServer, TaskJiraUsingToken,
		StoryJiraCloud, StoryJiraServer, StoryJiraUsingToken,
		BugJiraCloud, BugJiraServer, BugJiraUsingToken}
)
