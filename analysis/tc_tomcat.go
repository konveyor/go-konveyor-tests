package analysis

import (
	"github.com/konveyor/tackle2-hub/api"
)

var Tomcat = TC{
	Name: "Customer Tomcat Legacy - shoud never fail",
	Application: api.Application{
		Name: "Customer Tomcat Legacy",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor/example-applications.git",
			Path: "example-1",
		},
	},
	Task: Analyze,
	//ReportContent: map[string][]string{ // From Windup
	//	"/windup/report/index.html": {
	//		"1\nstory points",
	//		"1\nCloud Mandatory",
	//		"6\nInformation",
	//	},
	//},
	Analysis: api.Analysis{
		Effort: 2,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "Hardcoded IP Address\nWhen migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
				Effort: 1,
				RuleSet: "discovery-rules",
				Rule: "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File: "/addon/source/example-applications/example-1/target/classes/persistence.properties",
						Line: 0,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
					{
						File: "/addon/source/example-applications/example-1/src/main/resources/persistence.properties",
						Line: 0,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML"},
		{Name: "Java EE Batch"},
		{Name: "Properties"},
		{Name: "Servlet"},
	},
}
