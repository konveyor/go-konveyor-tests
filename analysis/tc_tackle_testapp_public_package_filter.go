package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestAppPackageFilter = api.Application{
	Name: "Tackle Testapp public",
	Repository: &api.Repository{
		Kind: "git",
		URL:  "https://github.com/konveyor/tackle-testapp-public",
	},
}

var TackleTestappPublicPackageFilter = TC{
	Name:        "Tackle Testapp public with package filter",
	Application: TackleTestApp,
	Task:        Analyze,
	WithDeps:    false,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=linux",
		},
	},
	Identities: []api.Identity{},
	Scope: &addon.Scope{
		Packages: struct {
			Included []string "json:\"included,omitempty\""
			Excluded []string "json:\"excluded,omitempty\""
		}{
			Included: []string{"com.example"},
		},
	},
	Analysis: api.Analysis{
		Effort: 1,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Hardcoded IP Address",
				Effort:      1,
				RuleSet:     "discovery-rules",
				Rule:        "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File:     "/addon/source/tackle-testapp-public/src/main/resources/persistence.properties",
						Line:     2,
						Message:  "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
						CodeSnip: "jdbc.url=jdbc:oracle:thin:@10.19.2.93:15",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
	},
}
