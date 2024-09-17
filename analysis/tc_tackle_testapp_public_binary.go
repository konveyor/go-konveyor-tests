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
	Analysis: api.Analysis{
		Effort: 1,
		Issues: []api.Issue{
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
	},
}
