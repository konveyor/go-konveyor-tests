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
	Analysis: api.Analysis{
		Effort: 5,
		Issues: []api.Issue{},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet"},
		{Name: "Java Threads"},
	},
}