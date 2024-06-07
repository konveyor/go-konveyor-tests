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
	Labels: api.Map{
		"included": []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Analysis: api.Analysis{
		Effort: 5,
		Issues: []api.Issue{
			{Category: "cloud-mandatory", Description: "Trivial change or 1-1 library swap", Effort: 2, Name: "File system - Java IO"},
			{Category: "cloud-mandatory", Description: "Trivial change or 1-1 library swap", Effort: 3, Name: "Hard-coded IP address"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Java Servlet"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Threads"},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet"},
		{Name: "Java Threads"},
	},
}
