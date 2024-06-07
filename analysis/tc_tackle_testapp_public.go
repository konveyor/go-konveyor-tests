package analysis

import (
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
	Name:        "Tackle Testapp public",
	Application: TackleTestApp,
	Task:        Analyze,
	WithDeps:    false,
	Labels: api.Map{
		"included": []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=linux",
		},
	},
	Identities: []api.Identity{},
	Analysis: api.Analysis{
		Effort: 8,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Hardcoded IP Address\nWhen migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
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
			{
				Category:    "mandatory",
				Description: "File system - Java IO\nAn application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				Effort:      1,
				RuleSet:     "openshift",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:     "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:     45,
						Message:  "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
						CodeSnip: "dataSource.setDriverClassName(config.getProperty(\"jdbc.driverClassName\"));",
					},
					{
						File:     "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:     46,
						Message:  "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
						CodeSnip: "dataSource.setUrl(config.getProperty(\"jdbc.url\"));",
					},
					{
						File:     "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:     47,
						Message:  "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
						CodeSnip: "dataSource.setUsername(config.getProperty(\"jdbc.user\"));",
					},
					{
						File:     "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:     48,
						Message:  "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
						CodeSnip: "dataSource.setPassword(config.getProperty(\"jdbc.password\"));",
					},
					{
						File:     "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:     56,
						Message:  "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
						CodeSnip: "transactionManager.setEntityManagerFactory(",
					},
					{
						File:     "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:     68,
						Message:  "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
						CodeSnip: "hibernateProperties.setProperty(\"hibernate.hbm2ddl.auto\", config.getProperty(\"hibernate.hbm2ddl.auto\"));",
					},
					{
						File:     "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:     69,
						Message:  "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
						CodeSnip: "hibernateProperties.setProperty(\"hibernate.dialect\", config.getProperty(\"hibernate.dialect\"));",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Processing"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Execute"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
	},
}
