package analysis

import "github.com/konveyor/tackle2-hub/api"

var PetclinicHazelcast = TC{
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
				Rule:        "hazelcast-cloud-readiness-hz001",
				Incidents: []api.Incident{
					{
						File: "SessionConfiguration_java.html",
						Line: 78,
					},
					{
						File: "SessionConfiguration_java.html",
						Line: 77,
					},
					{
						File: "pom_xml.html",
						Line: 215,
					},
				},
			},
			{
				Category:    "cloud-mandatory",
				Description: "Trivial change or 1-1 library swap",
				Effort:      1,
				Name:        "Embedded Hazelcast dependencies",
				Rule:        "hazelcast-cloud-readiness-hz002",
			},
			{
				Category:    "cloud-mandatory",
				Description: "Trivial change or 1-1 library swap",
				Effort:      5,
				Name:        "File system - Java IO",
				Rule:        "local-storage-00001",
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
}
