package analysis

import "github.com/konveyor/tackle2-hub/api"

var SeamBooking = TC{
	Name: "Seam booking",
	Application: api.Application{
		Name: "Seam booking 5.2",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/windup/windup.git",
			Path:   "test-files/seam-booking-5.2",
			Branch: "master",
		},
	},
	Task: Analyze,
	ReportContent: map[string][]string{
		"/windup/report/index.html": {
			"0\nstory points",
			"3\nInformation",
		},
	},
	Targets: []string{
		"cloud-readiness",
	},
	Analysis: api.Analysis{
		Effort: 0,
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "EJB", Category: api.Ref{Name: "Sustain"}},
		{Name: "JSF XML", Category: api.Ref{Name: "View"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "EJB", Category: api.Ref{Name: "Clustering"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Execute"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Web"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Processing"}},
	},
}
