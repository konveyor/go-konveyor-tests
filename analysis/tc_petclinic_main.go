package analysis

import "github.com/konveyor/tackle2-hub/api"

var PetclinicMain = TC{
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
	Targets: []string{
		"cloud-readiness",
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
}
