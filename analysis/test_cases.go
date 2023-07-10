package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

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
	},
	{
		Name: "Petclinic cloud-readiness with tagger main",
		Application: api.Application{
			Name:        "Petclinic",
			Description: "Spring framework app",
			Repository: &api.Repository{
				Kind:   "git",
				URL:    "https://github.com/konveyor/spring-framework-petclinic.git",
				Branch: "main",
			},
		},
		Task: Analyze,
		ReportContent: map[string][]string{
			"/windup/report/index.html": {
				"5\nstory points",
				"5\nCloud Mandatory",
				"4\nInformation",
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
					Name:        "Embedded Spring Data JPA",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Embedded framework - Spring MVC",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Maven POM (pom.xml)",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Spring JMX configuration detected",
				},
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
				Kind:   "git",
				URL:    "https://github.com/savitharaghunathan/spring-framework-petclinic.git",
				Branch: "legacy",
			},
		},
		CustomRules: []api.RuleSet{
			{
				Name:   "Hazelcast Java distributed session store ruleset.",
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
		Task: Analyze,
		ReportContent: map[string][]string{
			"/windup/report/index.html": {
				"12\nstory points",
				"8\nCloud Mandatory",
				"13\nInformation",
			},
		},
		Analysis: api.Analysis{
			Effort: 12,
			Issues: []api.Issue{
				{
					Category:    "cloud-mandatory",
					Description: "Complex change with documented solution",
					Effort:      6,
					Name:        "Embedded Hazelcast",
				},
				{
					Category:    "cloud-mandatory",
					Description: "Trivial change or 1-1 library swap",
					Effort:      1,
					Name:        "Embedded Hazelcast dependencies",
				},
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
					Name:        "Bean Validation",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Common Annotations",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Embedded Spring Data JPA",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Embedded framework - Spring MVC",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "JAXB",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Java Servlet",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Maven POM (pom.xml)",
				},
				{
					Category:    "information",
					Description: "Info",
					Effort:      0,
					Name:        "Spring JMX configuration detected",
				},
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

// Shared parameters.
var Addons = []string{
	"windup", // legacy windup analyzer
	// "analyzer",	// LSP analyzer
}

var Analyze = api.Task{
	State: "Ready", // Created / Ready
	Data:  defaultTaskData,
}

var defaultTaskData = addon.Data{
	Output: "/windup/report",
	Mode: addon.Mode{
		Artifact: "",
		Binary:   false,
		WithDeps: false,
		Diva:     true,
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
