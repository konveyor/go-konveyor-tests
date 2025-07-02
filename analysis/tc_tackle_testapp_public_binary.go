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
		Effort: 2,
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
						Message: `An application running inside a container could lose access to a file in local storage.`,
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
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Logging to file system",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "logging-0000",
				Incidents: []api.Incident{
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/RollingFileAppender.java",
						Line: 3,
						Message: `If the application writes logs to a file system, local log files may be lost if an instance terminates or restarts.

						Recommendations

						* Use a centralized log management system.
						* Log to standard output and allow the cloud platform to manage the logging.
						* Use shared storage for log files.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/RollingPolicy.java",
						Line: 3,
						Message: `If the application writes logs to a file system, local log files may be lost if an instance terminates or restarts.

						Recommendations

						* Use a centralized log management system.
						* Log to standard output and allow the cloud platform to manage the logging.
						* Use shared storage for log files.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/RollingPolicyBase.java",
						Line: 3,
						Message: `If the application writes logs to a file system, local log files may be lost if an instance terminates or restarts.

						Recommendations

						* Use a centralized log management system.
						* Log to standard output and allow the cloud platform to manage the logging.
						* Use shared storage for log files.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/DemultiplexingLogHandler.java",
						Line: 6,
						Message: `If the application writes logs to a file system, local log files may be lost if an instance terminates or restarts.

						Recommendations

						* Use a centralized log management system.
						* Log to standard output and allow the cloud platform to manage the logging.
						* Use shared storage for log files.`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File: "/shared/bin/java-project/src/main/java/antlr/ImportVocabTokenManager.java",
						Line: 16,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/ImportVocabTokenManager.java",
						Line: 18,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/ImportVocabTokenManager.java",
						Line: 27,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/PreservingFileWriter.java",
						Line: 16,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/PreservingFileWriter.java",
						Line: 19,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/PreservingFileWriter.java",
						Line: 32,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/PreservingFileWriter.java",
						Line: 45,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/PreservingFileWriter.java",
						Line: 46,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/PreservingFileWriter.java",
						Line: 85,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/PreservingFileWriter.java",
						Line: 86,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 95,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 96,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 125,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 126,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 296,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 302,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 310,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 345,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/Tool.java",
						Line: 347,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/build/ANTLR.java",
						Line: 55,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/build/ANTLR.java",
						Line: 69,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/antlr/preprocessor/Hierarchy.java",
						Line: 85,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/classic/jmx/JMXConfigurator.java",
						Line: 75,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/classic/jmx/JMXConfigurator.java",
						Line: 86,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/classic/util/ContextInitializer.java",
						Line: 79,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/FileAppender.java",
						Line: 107,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/joran/GenericConfigurator.java",
						Line: 74,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/joran/GenericConfigurator.java",
						Line: 85,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/joran/action/PropertyAction.java",
						Line: 34,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/joran/spi/ConfigurationWatchList.java",
						Line: 75,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/property/FileExistsPropertyDefiner.java",
						Line: 23,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/recovery/ResilientFileOutputStream.java",
						Line: 17,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/recovery/ResilientFileOutputStream.java",
						Line: 35,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/FixedWindowRollingPolicy.java",
						Line: 83,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/FixedWindowRollingPolicy.java",
						Line: 90,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/RollingFileAppender.java",
						Line: 152,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/RollingFileAppender.java",
						Line: 54,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/SizeAndTimeBasedFNATP.java",
						Line: 61,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/TimeBasedFileNamingAndTriggeringPolicyBase.java",
						Line: 47,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 113,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 121,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 131,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 132,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 42,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 52,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 62,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/Compressor.java",
						Line: 63,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/RenameUtil.java",
						Line: 17,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/RenameUtil.java",
						Line: 19,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/RenameUtil.java",
						Line: 68,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/SizeAndTimeBasedArchiveRemover.java",
						Line: 12,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/rolling/helper/TimeBasedArchiveRemover.java",
						Line: 50,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/status/ViewStatusMessagesServletBase.java",
						Line: 154,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/util/FileUtil.java",
						Line: 86,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/util/FileUtil.java",
						Line: 87,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/util/LocationUtil.java",
						Line: 34,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/core/JsonFactory.java",
						Line: 345,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/core/JsonFactory.java",
						Line: 442,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/core/TokenStreamFactory.java",
						Line: 89,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ObjectReader.java",
						Line: 1063,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/deser/std/FromStringDeserializer.java",
						Line: 207,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 32,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 34,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 44,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/model/impl/RuntimeBuiltinLeafInfoImpl.java",
						Line: 224,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/runtime/MarshallerImpl.java",
						Line: 146,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/sax/SystemIdResolver.java",
						Line: 65,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/stax/factory/StAXOutputFactory.java",
						Line: 56,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/stax/factory/StAXOutputFactory.java",
						Line: 81,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/tools/PrintTable.java",
						Line: 81,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/tools/TransformInputOutput.java",
						Line: 28,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/tools/TransformInputOutput.java",
						Line: 35,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/tools/TransformInputOutput.java",
						Line: 36,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/txw2/output/StreamSerializer.java",
						Line: 44,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/io/konveyor/demo/config/ApplicationConfiguration.java",
						Line: 14,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/CtClass.java",
						Line: 559,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/DirClassPath.java",
						Line: 22,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/DirClassPath.java",
						Line: 33,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/bytecode/ClassFilePrinter.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/bytecode/ConstInfo.java",
						Line: 36,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/bytecode/ConstPool.java",
						Line: 654,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/bytecode/StackMapTable.java",
						Line: 70,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/tools/Dump.java",
						Line: 17,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/tools/Dump.java",
						Line: 19,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/tools/web/Webserver.java",
						Line: 183,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/tools/web/Webserver.java",
						Line: 186,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/util/HotSwapAgent.java",
						Line: 94,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/util/HotSwapAgent.java",
						Line: 98,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/util/HotSwapAgent.java",
						Line: 114,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/util/proxy/FactoryHelper.java",
						Line: 92,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/activation/FileDataSource.java",
						Line: 21,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/activation/FileDataSource.java",
						Line: 25,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/activation/FileDataSource.java",
						Line: 29,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/activation/MailcapCommandMap.java",
						Line: 534,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/activation/MimetypesFileTypeMap.java",
						Line: 250,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/xml/bind/JAXB.java",
						Line: 126,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/xml/bind/JAXB.java",
						Line: 221,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/xml/bind/helpers/AbstractMarshallerImpl.java",
						Line: 44,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/xml/bind/helpers/AbstractUnmarshallerImpl.java",
						Line: 104,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1009,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1040,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1045,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1062,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1067,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1081,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1097,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1486,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1488,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1538,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1560,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1571,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 1768,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/build/Plugin.java",
						Line: 211,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 681,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 683,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 735,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 787,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 812,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 878,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 896,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 199,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 203,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 228,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 244,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 251,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 314,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 500,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/DynamicType.java",
						Line: 503,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/loading/ClassInjector.java",
						Line: 80,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/loading/ClassInjector.java",
						Line: 87,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/scaffold/TypeWriter.java",
						Line: 248,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/babelfish/TranslationManager.java",
						Line: 26,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/LogDecryptor.java",
						Line: 27,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/LogDecryptor.java",
						Line: 47,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/LogDecryptor.java",
						Line: 48,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecureLogHandler.java",
						Line: 37,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecureLogHandler.java",
						Line: 44,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecureLogHandler.java",
						Line: 65,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecureLogHandler.java",
						Line: 84,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 154,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 182,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 229,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 231,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 243,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 284,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 286,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredLoggerImpl.java",
						Line: 366,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredLoggerImpl.java",
						Line: 387,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredMemoryLogHandler.java",
						Line: 460,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/ClioSupport.java",
						Line: 58,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/HAManager.java",
						Line: 443,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/OracleLog.java",
						Line: 49,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/OracleParameterMetaDataParser.java",
						Line: 395,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/OraclePreparedStatement.java",
						Line: 500,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/OraclePreparedStatement.java",
						Line: 555,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/OracleSql.java",
						Line: 1515,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/PropertiesFileUtil.java",
						Line: 248,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/PropertiesFileUtil.java",
						Line: 262,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/T4CTTIrxd.java",
						Line: 791,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/OracleType.java",
						Line: 320,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/OracleType.java",
						Line: 327,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/OracleTypeADT.java",
						Line: 2321,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/OracleTypeCHAR.java",
						Line: 173,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/OracleTypeCHAR.java",
						Line: 239,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/OracleTypeCHAR.java",
						Line: 58,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/OracleTypeCHAR.java",
						Line: 81,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/PickleContext.java",
						Line: 135,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/PickleContext.java",
						Line: 164,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/PickleContext.java",
						Line: 193,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/oracore/TypeTreeElement.java",
						Line: 68,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/proxy/ClassGenerator.java",
						Line: 127,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/proxy/ProxyExport.java",
						Line: 42,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/jdbc/nl/NLParamParser.java",
						Line: 355,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/jdbc/nl/NLParamParser.java",
						Line: 373,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/jdbc/nl/NLParamParser.java",
						Line: 72,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/jdbc/nl/NLParamParser.java",
						Line: 77,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/jndi/JndiAttrs.java",
						Line: 134,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/nt/CustomSSLSocketFactory.java",
						Line: 283,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/nt/CustomSSLSocketFactory.java",
						Line: 500,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/resolver/EnvVariableResolver.java",
						Line: 43,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/resolver/TNSNamesNamingAdapter.java",
						Line: 108,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/resolver/TNSNamesNamingAdapter.java",
						Line: 119,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/resolver/TNSNamesNamingAdapter.java",
						Line: 121,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/resolver/TNSNamesNamingAdapter.java",
						Line: 151,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/resolver/TNSNamesNamingAdapter.java",
						Line: 76,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/resolver/TNSNamesNamingAdapter.java",
						Line: 83,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 109,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 123,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 144,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 145,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 158,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 173,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 227,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 35,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 75,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 78,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/ConverterArchive.java",
						Line: 80,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/STRUCT.java",
						Line: 328,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/STRUCT.java",
						Line: 334,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/TypeDescriptor.java",
						Line: 926,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/sql/TypeDescriptor.java",
						Line: 937,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogProcessor.java",
						Line: 136,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogProcessor.java",
						Line: 145,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogProcessor.java",
						Line: 155,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogReader.java",
						Line: 89,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogScanner.java",
						Line: 17,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogWriter.java",
						Line: 26,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogWriter.java",
						Line: 32,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/HdrHistogram/HistogramLogWriter.java",
						Line: 38,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/apache/logging/log4j/simple/SimpleLoggerContext.java",
						Line: 45,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/apache/logging/log4j/simple/SimpleLoggerContext.java",
						Line: 46,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/apache/logging/log4j/status/StatusData.java",
						Line: 77,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/apache/logging/log4j/util/LowLevelLogUtil.java",
						Line: 31,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/apache/logging/log4j/util/LowLevelLogUtil.java",
						Line: 35,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/expression/function/FileFunction.java",
						Line: 90,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/fulltext/FullTextLucene.java",
						Line: 174,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/fulltext/FullTextLucene.java",
						Line: 266,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/fulltext/FullTextLucene.java",
						Line: 282,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/fulltext/FullTextLucene.java",
						Line: 335,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/message/TraceSystem.java",
						Line: 215,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 30,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 34,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 316,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 48,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 59,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 63,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 558,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 567,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 570,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 67,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 590,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/server/TcpServerThread.java",
						Line: 234,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/server/web/WebApp.java",
						Line: 429,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/server/web/WebApp.java",
						Line: 865,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/FileLock.java",
						Line: 157,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/RecoverTester.java",
						Line: 50,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/async/FileAsync.java",
						Line: 39,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 425,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 435,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 439,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 149,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 168,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 172,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 196,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 232,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 239,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 275,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 304,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 309,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 313,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 317,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 322,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 329,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 359,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 391,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 400,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 55,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 61,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 89,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 90,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomapped/FileNioMapped.java",
						Line: 35,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/zip/FilePathZip.java",
						Line: 249,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/zip/FilePathZip.java",
						Line: 299,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/ConvertTraceFile.java",
						Line: 61,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/ConvertTraceFile.java",
						Line: 62,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Recover.java",
						Line: 224,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 114,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 137,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 139,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 213,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/CloseWatcher.java",
						Line: 43,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/Profiler.java",
						Line: 104,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SortedProperties.java",
						Line: 107,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 174,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 319,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/ThreadDeadlockDetector.java",
						Line: 59,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JandexAntTask.java",
						Line: 37,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JarIndexer.java",
						Line: 141,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JarIndexer.java",
						Line: 142,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JarIndexer.java",
						Line: 26,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JarIndexer.java",
						Line: 48,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JarIndexer.java",
						Line: 49,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JarIndexer.java",
						Line: 53,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/JarIndexer.java",
						Line: 59,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 92,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 116,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 128,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 201,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 216,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 69,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 85,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jboss/jandex/Main.java",
						Line: 89,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jvnet/staxex/Base64Data.java",
						Line: 295,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/jvnet/staxex/Base64Data.java",
						Line: 337,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "File system - java.net.URL/URI",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00002",
				Incidents: []api.Incident{
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/classic/util/ContextInitializer.java",
						Line: 73,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/joran/action/IncludeAction.java",
						Line: 93,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/util/LocationUtil.java",
						Line: 30,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/deser/std/FromStringDeserializer.java",
						Line: 209,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/deser/std/StdKeyDeserializer.java",
						Line: 204,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 38,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/model/impl/RuntimeBuiltinLeafInfoImpl.java",
						Line: 236,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/model/impl/RuntimeBuiltinLeafInfoImpl.java",
						Line: 251,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/model/impl/RuntimeBuiltinLeafInfoImpl.java",
						Line: 667,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/schemagen/XmlSchemaGenerator.java",
						Line: 453,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/schemagen/XmlSchemaGenerator.java",
						Line: 454,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/tools/TransformInputOutput.java",
						Line: 105,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/tools/TransformInputOutput.java",
						Line: 55,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/sun/xml/fastinfoset/tools/TransformInputOutput.java",
						Line: 72,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/ByteArrayClassPath.java",
						Line: 33,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/JarClassPath.java",
						Line: 74,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/URLClassPath.java",
						Line: 93,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javassist/tools/web/Viewer.java",
						Line: 85,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/xml/bind/JAXB.java",
						Line: 124,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/xml/bind/JAXB.java",
						Line: 219,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/javax/xml/bind/helpers/ValidationEventLocatorImpl.java",
						Line: 60,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 1221,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/ClassFileLocator.java",
						Line: 651,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/loading/ByteArrayClassLoader.java",
						Line: 527,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/loading/PackageDefinitionStrategy.java",
						Line: 125,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/net/bytebuddy/dynamic/loading/PackageDefinitionStrategy.java",
						Line: 135,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/CharCommonAccessor.java",
						Line: 424,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/JavaToJavaConverter.java",
						Line: 1240,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/OldUpdatableResultSet.java",
						Line: 2992,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/rowset/OracleCachedRowSet.java",
						Line: 2391,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/security/auth/DefaultAuthenticator.java",
						Line: 113,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 396,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Server.java",
						Line: 445,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 260,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/hibernate/validator/internal/constraintvalidators/hv/URLValidator.java",
						Line: 23,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "File system - Java NIO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00005",
				Incidents: []api.Incident{
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/FileAppender.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/FileAppender.java",
						Line: 8,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/ch/qos/logback/core/recovery/ResilientFileOutputStream.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/Java7HandlersImpl.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 12,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 13,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathDeserializer.java",
						Line: 14,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/com/fasterxml/jackson/databind/ext/NioPathSerializer.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 8,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/diagnostics/SecuredFileLogHandler.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/PropertiesFileUtil.java",
						Line: 8,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/PropertiesFileUtil.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/PropertiesFileUtil.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/PropertiesFileUtil.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/nt/CustomSSLSocketFactory.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/nt/CustomSSLSocketFactory.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/net/nt/CustomSSLSocketFactory.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/expression/function/FileFunction.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/expression/function/FileFunction.java",
						Line: 8,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/fulltext/FullTextLucene.java",
						Line: 4,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/DataUtils.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/FileStore.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/FileStore.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/MVStoreTool.java",
						Line: 8,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/cache/FilePathCache.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/cache/FilePathCache.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/mvstore/db/Store.java",
						Line: 4,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/server/web/WebServer.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/server/web/WebServer.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/server/web/WebServer.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/server/web/WebServer.java",
						Line: 12,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/FileLister.java",
						Line: 3,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/FileLock.java",
						Line: 12,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/FileLock.java",
						Line: 13,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/FileStore.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FakeFileChannel.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FakeFileChannel.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FileBase.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FileBase.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FileBaseDefault.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FileChannelInputStream.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FilePath.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FilePathWrapper.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FileUtils.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FileUtils.java",
						Line: 13,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/FileUtils.java",
						Line: 14,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/async/FileAsync.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/async/FileAsync.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/async/FileAsync.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/async/FilePathAsync.java",
						Line: 4,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 8,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 12,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 13,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 14,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 15,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 16,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 17,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 18,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 19,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 20,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 21,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 22,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 23,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/disk/FilePathDisk.java",
						Line: 24,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/encrypt/FileEncrypt.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/encrypt/FileEncrypt.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/encrypt/FilePathEncrypt.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/mem/FileMem.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/mem/FilePathMem.java",
						Line: 3,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomapped/FileNioMapped.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomapped/FileNioMapped.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomapped/FileNioMapped.java",
						Line: 12,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomapped/FileNioMapped.java",
						Line: 13,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomapped/FilePathNioMapped.java",
						Line: 4,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomem/FileNioMem.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/niomem/FilePathNioMem.java",
						Line: 3,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/rec/FilePathRec.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/rec/FileRec.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/rec/FileRec.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/retry/FilePathRetryOnInterrupt.java",
						Line: 4,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/retry/FileRetryOnInterrupt.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/retry/FileRetryOnInterrupt.java",
						Line: 8,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/split/FilePathSplit.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/split/FileSplit.java",
						Line: 5,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/split/FileSplit.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/zip/FilePathZip.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/zip/FileZip.java",
						Line: 6,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/store/fs/zip/FileZip.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/ChangeFileEncryption.java",
						Line: 7,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 9,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 10,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 11,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 12,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 13,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/Profiler.java",
						Line: 13,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/Profiler.java",
						Line: 14,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/Profiler.java",
						Line: 15,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 17,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 18,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 19,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 20,
						Message: `An application running inside a container could lose access to a file in local storage.

						Recommendations

						The following recommendations depend on the function of the file in local storage:

						* Logging: Log to standard output and use a centralized log collector to analyze the logs.
						* Caching: Use a cache backing service.
						* Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
						* Data storage: Use a database backing service for relational data or use a persistent data storage system.
						* Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Local JDBC Calls",
				Effort:      7,
				RuleSet:     "cloud-readiness",
				Rule:        "localhost-jdbc-00002",
				Incidents: []api.Incident{
					{
						File:    "/shared/bin/java-project/src/main/java/org/h2/server/web/WebServer.java",
						Line:    47,
						Message: `The app is trying to access local resource by JDBC, please try to migrate the resource to cloud`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Java native libraries (JNI, JNA)",
				Effort:      7,
				RuleSet:     "cloud-readiness",
				Rule:        "jni-native-code-00000",
				Incidents: []api.Incident{
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/T2CConnection.java",
						Line: 3114,
						Message: `Java native libraries might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native library in your application.
						* Check whether the native library is compatible with a cloud environment.
						* Reuse or embed the native library or application in a cloud environment, for example, in a JBoss module.
						* Replace, remove, or rewrite the native library or application using a cloud-compatible equivalent.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/oracle/jdbc/xa/client/OracleXADataSource.java",
						Line: 365,
						Message: `Java native libraries might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native library in your application.
						* Check whether the native library is compatible with a cloud environment.
						* Reuse or embed the native library or application in a cloud environment, for example, in a JBoss module.
						* Replace, remove, or rewrite the native library or application using a cloud-compatible equivalent.`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Java Native Processes",
				Effort:      7,
				RuleSet:     "cloud-readiness",
				Rule:        "jni-native-code-00001",
				Incidents: []api.Incident{
					{
						File: "/shared/bin/java-project/src/main/java/antlr/build/Tool.java",
						Line: 69,
						Message: `Native Processes might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native process in your application.
						* Check whether the native process is compatible with a cloud environment.
						* Replace, remove, or rewrite the native process or application using a cloud-compatible equivalent.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 244,
						Message: `Native Processes might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native process in your application.
						* Check whether the native process is compatible with a cloud environment.
						* Replace, remove, or rewrite the native process or application using a cloud-compatible equivalent.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/tools/Upgrade.java",
						Line: 247,
						Message: `Native Processes might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native process in your application.
						* Check whether the native process is compatible with a cloud environment.
						* Replace, remove, or rewrite the native process or application using a cloud-compatible equivalent.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/Profiler.java",
						Line: 248,
						Message: `Native Processes might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native process in your application.
						* Check whether the native process is compatible with a cloud environment.
						* Replace, remove, or rewrite the native process or application using a cloud-compatible equivalent.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 291,
						Message: `Native Processes might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native process in your application.
						* Check whether the native process is compatible with a cloud environment.
						* Replace, remove, or rewrite the native process or application using a cloud-compatible equivalent.`,
					},
					{
						File: "/shared/bin/java-project/src/main/java/org/h2/util/SourceCompiler.java",
						Line: 294,
						Message: `Native Processes might not run in a cloud or container environment.

						Recommendations

						* Review the purpose of the native process in your application.
						* Check whether the native process is compatible with a cloud environment.
						* Replace, remove, or rewrite the native process or application using a cloud-compatible equivalent.`,
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
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_ar.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_cs.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_da.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_de.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_el.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_es.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_fi.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_fr.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_hu.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_it.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_iw.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_ja.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_ko.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_nl.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_no.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_pl.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_pt.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_pt_BR.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_ro.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_ru.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_sk.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_sv.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_th.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_tr.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_zh_CN.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/ojdbc11-21-1-0-0-jar-exploded/oracle/jdbc/driver/Messages_zh_TW.properties",
						Line:    918,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/sun/xml/bind/v2/schemagen/XmlSchemaGenerator.java",
						Line:    856,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/oracle/jdbc/OracleDatabaseMetaData.java",
						Line:    27,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/BuildInfo.java",
						Line:    34,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/oracle/jdbc/driver/T2CConnection.java",
						Line:    59,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/oracle/net/ano/AuthenticationService.java",
						Line:    217,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/oracle/net/ano/AuthenticationService.java",
						Line:    218,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/org/h2/server/web/WebThread.java",
						Line:    236,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
					},
					{
						File:    "/shared/bin/java-project/src/main/java/org/h2/util/NetUtils.java",
						Line:    202,
						Message: `When migrating environments, hard-coded IP addresses may need to be modified or eliminated.`,
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
