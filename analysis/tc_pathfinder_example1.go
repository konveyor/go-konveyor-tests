package analysis

import "github.com/konveyor/tackle2-hub/api"

var PathfinderExample1 = TC{
	Name: "Pathfinder example1 cloud-readiness with tagger",
	Application: api.Application{
		Name:        "Pathfinder example-1",
		Description: "Tackle Pathfinder application.",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/konveyor/tackle-pathfinder.git",
			Branch: "1.2.0",
		},
	},
	Task: Analyze,
	//ReportContent: map[string][]string{  // From Windup
	//	"/windup/report/index.html": {
	//		"5\nstory points",
	//		"5\nCloud Mandatory",
	//		"9\nInformation",
	//	},
	//},
	Targets: []string{
		"cloud-readiness",
	},
	Analysis: api.Analysis{
		Effort: 271,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "File system - Java IO\nAn application running inside a container could lose access to a file in local storage..",
				Effort: 1,
				RuleSet: "openshift",
				Rule: "local-storage-00001",
				Incidents: []api.Incident{},	// Got 271 incidents for this issue.
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "CDI"},
		{Name: "Properties"},
		{Name: "Bean Validation"},
		{Name: "Application Properties File"},
	},
}
