package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPublicPackageFilter = TC{
	Name:        "Tackle Testapp public with package filter",
	Application: data.TackleTestappPublic,
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
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "Hardcoded IP Address",
				Effort:      1,
				RuleSet:     "discovery-rules",
				Rule:        "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File:     "/shared/source/tackle-testapp-public/src/main/resources/persistence.properties",
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
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
	},
}
