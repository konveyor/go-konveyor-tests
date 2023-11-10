package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/api/identity"
)

var MavenPublic = identity.Mvn

var TackleTestappPublicWithDeps = TC{
	Name:        "Tackle Testapp public with deps",
	Application: TackleTestApp,
	Task:        Analyze,
	WithDeps:    true,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=linux",
			"konveyor.io/target=cloud-readiness",
		},
	},
	Identities: []api.Identity{
		MavenPublic,	// Tackle Testapp public Maven registry expects MAVEN_TESTAPP_USER and MAVEN_TESTAPP_TOKEN env variables.
	},
	Analysis: api.Analysis{
		Effort: 9,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Hardcoded IP Address\nWhen migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
				Effort:      1,
				RuleSet:     "discovery-rules",
				Rule:        "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File:    "/addon/source/tackle-testapp-public/target/classes/persistence.properties",
						Line:    2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
					{
						File:    "/addon/source/tackle-testapp-public/src/main/resources/persistence.properties",
						Line:    2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
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
						File:    "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    45,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File:    "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    46,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File:    "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    47,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File:    "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    48,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File:    "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    56,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File:    "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    68,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
					{
						File:    "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    69,
						Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "org.apache.tomcat.tomcat-servlet-api",
				Version:  "9.0.46",
				Provider: "",
			},
			{
				Name:     "org.springframework.spring-jdbc",
				Version:  "5.3.7",
				Provider: "",
			},
			{
				Name:     "org.springframework.spring-webmvc",
				Version:  "5.3.7",
				Provider: "",
			},
			{
				Name:     "org.springframework.spring-web",
				Version:  "5.3.7",
				Provider: "",
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter-actuator",
				Version:  "2.5.0",
				Provider: "",
			},
			{
				Name:     "org.apache.tomcat.tomcat-jdbc",
				Version:  "9.0.46",
				Provider: "",
			},
			{
				Name:     "org.hibernate.hibernate-entitymanager",
				Version:  "5.4.32.Final",
				Provider: "",
			},
			{
				Name:     "org.hibernate.validator.hibernate-validator",
				Version:  "6.2.0.Final",
				Provider: "",
			},
			{
				Name:     "ch.qos.logback.logback-classic",
				Version:  "1.1.7",
				Provider: "",
			},
			{
				Name:     "com.oracle.database.jdbc.ojdbc11",
				Version:  "21.1.0.0",
				Provider: "",
			},
			{
				Name:     "org.postgresql.postgresql",
				Version:  "42.2.23",
				Provider: "",
			},
			{
				Name:     "com.h2database.h2",
				Version:  "2.1.214",
				Provider: "",
			},
			{
				Name:     "io.konveyor.demo.configuration-utils",
				Version:  "1.0.0",
				Provider: "",
			},
			{
				Name:     "com.fasterxml.jackson.jackson-bom",
				Version:  "2.12.3",
				Provider: "",
			},
			{
				Name:     "org.springframework.data.spring-data-bom",
				Version:  "2021.0.1",
				Provider: "",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Processing"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Web"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Execute"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Spring Web", Category: api.Ref{Name: "View"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Embedded"}},
	},
}
