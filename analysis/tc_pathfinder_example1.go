package analysis

import "github.com/konveyor/tackle2-hub/api"

var PathfinderExample1 = TC{
	Name: "Pathfinder example1 cloud-readiness with tagger",
	Application: api.Application{
		Name:        "Pathfinder example-1",
		Description: "Tackle Pathfinder application.",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/konveyor/tackle-pathfinder.git",
			Branch: "1.2.0",
		},
	},
	Task: Analyze,
	//ReportContent: map[string][]string{  // From Windup
	//	"/windup/report/index.html": {
	//		"5\nstory points",
	//		"5\nCloud Mandatory",
	//		"9\nInformation",
	//	},
	//},
	Targets: []string{
		"cloud-readiness",
	},
	Analysis: api.Analysis{
		Effort: 271,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "File system - Java IO\nAn application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				Effort: 1,
				RuleSet: "openshift",
				Rule: "local-storage-00001",
				Incidents: []api.Incident{
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/AssessmentsResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/TranslatorResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/controllers/TranslatorResource.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/main/java/io/tackle/pathfinder/services/AssessmentSvc.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/AssessmentsResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/controllers/TranslatorResourceTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/AssessmentSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				{
					File: "/addon/source/tackle-pathfinder/src/test/java/io/tackle/pathfinder/services/TranslatorSvcTest.java",
					Line: 0,
					Message: "An application running inside a container could lose access to a file in local storage.. Recommendations. The following recommendations depend on the function of the file in local storage:. * Logging: Log to standard output and use a centralized log collector to analyze the logs.. * Caching: Use a cache backing service.. * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.. * Data storage: Use a database backing service for relational data or use a persistent data storage system.. * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.",
				},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "CDI"},
		{Name: "Properties"},
		{Name: "Bean Validation"},
		{Name: "Application Properties File"},
	},
}
