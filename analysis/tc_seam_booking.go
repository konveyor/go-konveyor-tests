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
	Labels: api.Map{
		"included": []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Analysis: api.Analysis{
		Effort: 0,
		Issues: []api.Issue{
			{Category: "information", Description: "Info", Effort: 0, Name: "EJB XML"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Maven POM (pom.xml)"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Web XML"},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML"},
		{Name: "JBoss EJB XML"},
		{Name: "Stateful (SFSB)"},
		{Name: "Stateless (SLSB)"},
		{Name: "JDBC XA datasources"},
		{Name: "CSS"},
		{Name: "HTML"},
		{Name: "Properties"},
		{Name: "JPA entities"},
		{Name: "JPA named queries"},
		{Name: "JPA XML"},
		{Name: "Persistence units"},
		{Name: "JTA"},
		{Name: "JSF Page"},
		{Name: "Web XML File"},
		{Name: "JSF XML"},
		{Name: "EJB"},
		{Name: "Java EE XML"},
	},
}
