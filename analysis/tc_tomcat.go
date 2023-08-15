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
	ReportContent: map[string][]string{
		"/windup/report/index.html": {
			"1\nstory points",
			"1\nCloud Mandatory",
			"6\nInformation",
		},
	},
	Analysis: api.Analysis{
		Effort: 2,
		Issues: []api.Issue{
			{Category: "cloud-mandatory", Description: "Trivial change or 1-1 library swap", Effort: 1, Name: "Hard-coded IP address"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Embedded Spring Data JPA"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Embedded framework - Spring MVC"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Embedded framework - Spring Web"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Embedded library - Spring Boot Actuator"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Java Servlet"},
			{Category: "information", Description: "Info", Effort: 0, Name: "Maven POM (pom.xml)"},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "JPA entities"},
		{Name: "Properties"},
		{Name: "Servlet"},
		{Name: "Spring Boot Actuator"},
		{Name: "Spring Data JPA"},
		{Name: "Spring MVC"},
		{Name: "Spring Web"},
	},
}
