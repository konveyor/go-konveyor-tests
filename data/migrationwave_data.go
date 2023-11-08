package data

type ExportApplicationsCase struct {
	JiraInstance  JiraInstanceTC
	NumOfApps     int
	TicketKind    string
	TicketParent  string
}
