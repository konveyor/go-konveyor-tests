package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/data/identity"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPublicBinary = TC{
	Name:        "tackle-testapp-binary",
	Application: data.TackleTestappPublicBinary,
	Identities: []api.Identity{
		identity.TackleTestappPublicMaven,
	},
	Task: Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Binary: true,
	WithDeps: true,
	Analysis: api.Analysis{
		Effort: 2,
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/bin/java-project/src/main/java/io/konveyor/demo/config/ApplicationConfiguration.java",
						Line:    14,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Hardcoded IP Address",
				Effort:      1,
				RuleSet:     "discovery-rules",
				Rule:        "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File:    "/shared/bin/java-project/customers-tomcat-0-0-1-20240913-093117-1-war-exploded/WEB-INF/classes/persistence.properties",
						Line:    2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
	},
}
