package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

//
// Test cases for Application Analysis.
var TestCases = []TC{
	{
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
		Task:     Windup,
		ReportContent: map[string][]string{
			"/windup/report/index.html": {
				"5\nstory points",
				"5\nCloud Mandatory",
				"9\nInformation",
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
	},
	{
		Name: "Petclinic cloud-readiness with tagger main",
		Application: api.Application{
			Name:        "Petclinic",
			Description: "Spring framework app",
			Repository: &api.Repository{
				Kind: "git",
				URL:  "https://github.com/savitharaghunathan/spring-framework-petclinic.git",
				Branch: "main",
			},
		},
		Task:     Windup,
		ReportContent: map[string][]string{
			"/windup/report/index.html": {
				"5\nstory points",
				"5\nCloud Mandatory",
				"4\nInformation",
			},
		},
		AnalysisTags: []api.Tag{
			{Name: "Spring MVC"},
			{Name: "CSS"},
			{Name: "Properties"},
			{Name: "Spring JMX"},
			{Name: "JPA entities"},
			{Name: "Spring Data JPA"},
			{Name: "JSP Page"},
		},
	},
	{
		Name: "Petclinic legacy cloud-readiness with tagger and hazelcast custom rules",
		Application: api.Application{
			Name:        "Petclinic",
			Description: "Spring framework app",
			Repository: &api.Repository{
				Kind: "git",
				URL:  "https://github.com/savitharaghunathan/spring-framework-petclinic.git",
				Branch: "legacy",
			},
		},
		CustomRules: []api.RuleSet{
			{
				Name: "Hazelcast Java distributed session store ruleset.",
				Custom: true,
				Image: api.Ref{
					ID: 1,
				},
				Rules: []api.Rule{
					{
						File: &api.Ref{
							Name: "./data/hz.windup.xml",
						},
					},
				},
			},
		},
		Task:     Windup,
		ReportContent: map[string][]string{
			"/windup/report/index.html": {
				"12\nstory points",
				"8\nCloud Mandatory",
				"13\nInformation",
			},
		},
		AnalysisTags: []api.Tag{
			{Name: "Java EE JAXB"},
			{Name: "Servlet"},
			{Name: "Spring MVC"},
			{Name: "Spring JMX"},
			{Name: "Common Annotations"},
			{Name: "Properties"},
			{Name: "JPA entities"},
			{Name: "Spring Data JPA"},
			{Name: "Bean Validation"},
			{Name: "JSP Page"},
		},
	},
}

//
// Shared parameters.
var Addons = []string{
	"windup",	// legacy windup analyzer
//	"analyzer",	// LSP analyzer
}

var Windup = api.Task{
	State: "Ready",	// Created / Ready
	Data: defaultTaskData,
}

var defaultTaskData = addon.Data{
	Output: "/windup/report",
	Mode: addon.Mode{
		Artifact: "",
		Binary: false,
		WithDeps: false,
		Diva: true,
	},
	Sources: []string{},
	Targets: []string{"cloud-readiness"},
	Scope: addon.Scope{
		WithKnown: false,
		//Packages: {
		//	Included: []string{},
		//	Excluded: []string{},
		//},
	},
	Rules: addon.Rules{
		Path: "",
		Labels: []string{
			"cloud-readiness",
		},
		//Tags: {
		//	Excluded: []string{},
		//},
	},
	Tagger: addon.Tagger{
		Enabled: true,
	},
}
