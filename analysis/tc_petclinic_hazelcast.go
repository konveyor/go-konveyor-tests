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
	Targets: []string{
		"cloud-readiness",
	},
	Analysis: api.Analysis{
		Effort: 14,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "Local JDBC Calls",
				Effort: 7,
				RuleSet: "cloud-readiness",
				Rule: "localhost-jdbc-00002",
				Incidents: []api.Incident{
					{
						File: "/shared/source/spring-framework-petclinic/pom.xml",
						Line: 560,
						Message: "The app is trying to access local resource by JDBC, please try to migrate the resource to cloud",
					},
					{
						File: "/shared/source/spring-framework-petclinic/pom.xml",
						Line: 579,
						Message: "The app is trying to access local resource by JDBC, please try to migrate the resource to cloud",
					},
				},
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
