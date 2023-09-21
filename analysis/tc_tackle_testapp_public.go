package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestApp = api.Application{
	Name: "Tackle Testapp public",
	Repository: &api.Repository{
		Kind: "git",
		URL:  "https://github.com/konveyor/tackle-testapp-public",
	},
}

var TackleTestappPublic = TC{
	Name: "Tackle Testapp public",
	Application: TackleTestApp,
	Task: Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=linux",
			"konveyor.io/target=cloud-readiness",
		},
	},
	Analysis: api.Analysis{
		Effort: 10,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "Hardcoded IP Address\nWhen migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
				Effort: 1,
				RuleSet: "discovery-rules",
				Rule: "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/target/classes/persistence.properties",
						Line: 2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/resources/persistence.properties",
						Line: 2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "File system - Java IO\nAn application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				Effort: 1,
				RuleSet: "openshift",
				Rule: "local-storage-00001",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line: 44,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line: 45,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line: 46,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line: 47,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line: 55,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line: 67,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line: 68,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Windows file system path\nThis file system path is Microsoft Windows platform dependent. It needs to be replaced with a Linux-style path.",
				Effort:      1,
				RuleSet:     "os/windows",
				Rule:        "os-specific-00001",
				Incidents: []api.Incident{
					{
						File:    "/addon/source/tackle-testapp-public/.git/hooks/fsmonitor-watchman.sample",
						Line:    0,	// Fix.
						Message: "This file system path is Microsoft Windows platform dependent. It needs to be replaced with a Linux-style path.",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		// Fix.
	},
}
