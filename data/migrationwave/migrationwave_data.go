package migrationwave

import (
	"github.com/konveyor/go-konveyor-tests/data/jira"
)

type ExportApplicationsCase struct {
	JiraInstance jira.JiraInstanceTC
	NumOfApps    int
	TicketKind   string
	TicketParent string
}
