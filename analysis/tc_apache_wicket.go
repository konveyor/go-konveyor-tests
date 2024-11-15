package analysis

import "github.com/konveyor/tackle2-hub/api"

var ApacheWicket = TC{
	Name: "Apache Wicket",
	Application: api.Application{
		Name: "Apache Wicket app",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/windup/windup-sample-apps.git",
			Path: "test-files/src_example/org/apache/wicket",
		},
	},
	Task: Analyze,
	ReportContent: map[string][]string{
		"/windup/report/index.html": {
			"5\nstory points",
		},
	},
	Targets: []string{
		"cloud-readiness",
	},
	Analysis: api.Analysis{
		Effort: 10,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Effort: 1,
				Description: "Hardcoded IP Address",
				RuleSet: "discovery-rules",
				Rule: "hardcoded-ip-address",
			},
			{
				Category: "mandatory",
				Effort: 7,
				Description: "Local HTTP Calls",
				RuleSet: "cloud-readiness",
				Rule: "localhost-http-00001",
			},
		},
	},
}
