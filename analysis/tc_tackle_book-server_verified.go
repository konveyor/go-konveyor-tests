package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var BookServerApp = api.Application{
	Name: "BookServer App",
	Repository: &api.Repository{
		Kind: "git",
		URL:  "https://github.com/ibraginsky/book-server",
	},
}

var BookServerVerified = TC{
	Name:        "BookServer verified results",
	Application: BookServerApp,
	Task:        Analyze,
	WithDeps:    false,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=linux",
			"konveyor.io/target=openjdk21",
		},
	},
	Analysis: api.Analysis{
		Effort: 9,	// 6 without deps
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/controller/BookStartEndController.java",
						Line:    18,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/controller/BookStartEndController.java",
						Line:    24,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/model/BookModel.java",
						Line:    23,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/model/BookModel.java",
						Line:    40,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
				},
			},
			// Following comes with WithDeps: true
			{
				Category:    "mandatory",
				Description: "The Lombok version is incompatible with Open JDK 17",
				Effort:      3,
				RuleSet:     "openjdk17/openjdk11",
				Rule:        "lombok-incompatibility-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/book-server/pom.xml",
						Line:    31,
						Message: "Lombok supports Java 17 since version 1.18.22. The version of Lombok used in this project is too old and not compatible with Java 17. You should consider upgrading it.",
					},
				},
			},
		},
	},
}
