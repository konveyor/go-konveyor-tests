package analysis

import "github.com/konveyor/tackle2-hub/api"

var Daytrader = TC{
	Name: "Daytrader",
	Application: api.Application{
		Name: "Daytrader 7 EE application",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/WASdev/sample.daytrader7.git",
		},
	},
	WithDeps: true,
	Task:     Analyze,
	ReportContent: map[string][]string{
		"/windup/report/index.html": {
			"0\nstory points",
			"6\nInformation",
		},
	},
	Targets: []string{
		"cloud-readiness",
		"quarkus",
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
		{Name: "JDBC XA datasources"},
		{Name: "CSS"},
		{Name: "HTML"},
		{Name: "Properties"},
		{Name: "JPA entities"},
		{Name: "JPA XML"},
		{Name: "Persistence units"},
		{Name: "JTA"},
		{Name: "JSF Page"},
		{Name: "JSP Page"},
		{Name: "Web XML File"},
		{Name: "CDI XML"},
		{Name: "JSF XML"},
	},
}
