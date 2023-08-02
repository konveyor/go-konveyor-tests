package analysis

import "github.com/konveyor/tackle2-hub/api"

var PathfinderSample = TC{
	Name: "Pathfinder cloud-readiness with tagger",
	Application: api.Application{
		Name:        "Pathfinder",
		Description: "Tackle Pathfinder application.",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/konveyor/tackle-pathfinder.git",
			Branch: "1.2.0",
		},
	},
	Task: Analyze,
	ReportContent: map[string][]string{
		"/windup/report/index.html": {
			"5\nstory points",
			"5\nCloud Mandatory",
			"9\nInformation",
		},
	},
	Analysis: api.Analysis{
		Effort: 5,
		Issues: []api.Issue{
			{
				Category:    "cloud-mandatory",
				Description: "Trivial change or 1-1 library swap",
				Effort:      5,
				Name:        "File system - Java IO",
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				Name:        "Application properties file detected",
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				Name:        "Bean Validation",
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				Name:        "Maven POM (pom.xml)",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "CDI"},
		{Name: "HTML"},
		{Name: "Properties"},
		{Name: "JPA entities"},
		{Name: "JAX-RS"},
		{Name: "Bean Validation"},
		{Name: "Application Properties File"},
	},
}
