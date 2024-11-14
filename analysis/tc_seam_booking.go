package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var SeamBooking = TC{
	Name: "Seam booking",
	Application: api.Application{
		Name: "Seam booking 5.2",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/windup/windup.git",
			Path:   "test-files/seam-booking-5.2",
			Branch: "master",
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
			"0\nstory points",
			"3\nInformation",
		},
	},
	Analysis: api.Analysis{
		Effort: 3,
		Issues: []api.Issue{
			{
				Category:    "potential",
				Description: "The groupId `javax` has been replaced by `jakarta` in JBoss EAP 7.3, or later",
				Effort:       1,
				RuleSet:     "eap7/weblogic/tests/data",
				Rule:        "maven-javax-to-jakarta-00004",
				Name:        "Seam booking 5.2 JLYDS",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/pom.xml",
						Line:    24,
						Message: "If you migrate your application to JBoss EAP 7.3, or later, and want to ensure its Maven building, running or testing works as expected, use instead the Jakarta EE dependency - groupId jakarta..",
					},
				},
			},
			{
				Category: "potential",
				Description: "web.xml element references a javax-prefixed class name",
				Effort: 1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-servlet-00130",
				Name: "Maven POM (pom.xml)",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/resources/WEB-INF/web.xml",
						Line:    40,
						Message: "web.xml element references a javax-prefixed class name",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/resources/WEB-INF/web.xml",
						Line:    51,
						Message: "web.xml element references a javax-prefixed class name",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "EJB", Category: api.Ref{Name: "Sustain"}},
		{Name: "JSF XML", Category: api.Ref{Name: "View"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "EJB", Category: api.Ref{Name: "Clustering"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Execute"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Web"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Processing"}},
	},
}
