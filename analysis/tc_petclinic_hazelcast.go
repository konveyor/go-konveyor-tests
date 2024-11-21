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
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=jakarta-ee",
		},
	},
	ReportContent: map[string][]string{
		"/windup/report/index.html": {
			"12\nstory points",
			"8\nCloud Mandatory",
			"13\nInformation",
		},
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
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Validation"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Store"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
	},
}
