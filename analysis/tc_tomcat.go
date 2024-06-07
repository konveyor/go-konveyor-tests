package analysis

import (
	"github.com/konveyor/tackle2-hub/api"
)

var Tomcat = TC{
	Name: "Customer Tomcat Legacy - shoud never fail",
	Application: api.Application{
		Name: "Customer Tomcat Legacy",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor/example-applications.git",
			Path: "example-1",
		},
	},
	Task: Analyze,
	Labels: api.Map{
		"included": []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=linux",
		},
	},
	Analysis: api.Analysis{
		Effort: 1,
		Issues: []api.Issue{
			{
				Category:    "cloud-mandatory",
				Description: "Trivial change or 1-1 library swap",
				Effort:      1,
				RuleSet:     "DiscoverHardcodedIPAddressRuleProvider",
				Rule:        "DiscoverHardcodedIPAddressRuleProvider",
				Incidents: []api.Incident{
					{
						File:    "persistence.properties",
						Line:    2,
						Message: "Hard-coded IP: 169.60.225.216When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "database",
				Rule:        "database-03200",
				Incidents: []api.Incident{
					{
						File:    "pom.xml",
						Line:    0,
						Message: "The application embeds Spring Data JPA",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "mvc",
				Rule:        "mvc-01220",
				Incidents: []api.Incident{
					{
						File:    "pom.xml",
						Line:    0,
						Message: "The application embeds a Spring MVC library.",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "embedded-framework",
				Rule:        "embedded-framework-08400",
				Incidents: []api.Incident{
					{
						File:    "pom.xml",
						Line:    0,
						Message: "The application embeds the Spring Web framework.",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "observability",
				Rule:        "observability-0100",
				Incidents: []api.Incident{
					{
						File:    "pom.xml",
						Line:    0,
						Message: "The application embeds Spring Boot Actuator.",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "technology-usage-javaee",
				Rule:        "javaee-technology-usage-00120",
				Incidents: []api.Incident{
					{
						File:    "OrderManagementAppInitializer.java",
						Line:    0,
						Message: "The application uses Java Servlets",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "DiscoverMavenProjectsRuleProvider",
				Rule:        "DiscoverMavenProjectsRuleProvider_1",
				Incidents: []api.Incident{
					{
						File:    "pom.xml",
						Line:    0,
						Message: "Maven Project Object Model (POM) File",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet"},
		{Name: "Spring MVC"},
		{Name: "Spring Boot Actuator"},
		{Name: "Properties"},
		{Name: "JPA entities"},
		{Name: "Spring Data JPA"},
		{Name: "Spring Web"},
	},
}
