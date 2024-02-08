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
	ExportTaskToJiraCloud = ExportApplicationsCase{
		JiraInstance: jira.JiraCloud,
		TicketKind:   "10007", /* Task issuetypeId */
		TicketParent: "10001"} /* mta_integration projectId */

	ExportBugToJiraCloud = ExportApplicationsCase{
		JiraInstance: jira.JiraCloud,
		TicketKind:   "10012", /* Bug */
		TicketParent: "10001"}

	ExportTaskToJiraServer = ExportApplicationsCase{
		JiraInstance: jira.JiraServer,
		TicketKind:   "3",        /* Task issuetypeId */
		TicketParent: "12340621"} /* mta-qe-test projectId */

	ExportStoryToJiraServer = ExportApplicationsCase{
		JiraInstance: jira.JiraServer,
		TicketKind:   "17", /* Story */
		TicketParent: "12340621"}
	ExportBugToJiraServer = ExportApplicationsCase{
		JiraInstance: jira.JiraServer,
		TicketKind:   "1", /* Bug */
		TicketParent: "12340621"}

	ExportTaskToJiraUsingToken = ExportApplicationsCase{
		JiraInstance: jira.JiraServerBearerToken, /* Using token for Jira connection */
		TicketKind:   "3",                        /* Task issuetypeId */
		TicketParent: "12340621"}                 /* mta-qe-test projectId */
	ExportStoryToJiraUsingToken = ExportApplicationsCase{
		JiraInstance: jira.JiraServerBearerToken,
		TicketKind:   "17", /* Story */
		TicketParent: "12340621"}
	ExportBugToJiraUsingToken = ExportApplicationsCase{
		JiraInstance: jira.JiraServerBearerToken,
		TicketKind:   "1", /* Bug */
		TicketParent: "12340621"}
)
