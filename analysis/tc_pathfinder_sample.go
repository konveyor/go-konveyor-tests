package analysis

import "github.com/konveyor/tackle2-hub/api"

var PathfinderSample = TC{
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
	Targets: []string{
		"cloud-readiness",
	},
	Analysis: api.Analysis{
		Effort: 5,
		Issues: []api.Issue{
			{
				Category:    "cloud-mandatory",
				Description: "Trivial change or 1-1 library swap",
				Effort:      5,
				RuleSet:     "local-storage",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:    "MavenWrapperDownloader.java",
						Line:    111,
						Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nTwelve-Factor App: Backing servicesTwelve-Factor App: ConfigOpenShift Container Platform: Input secrets and ConfigMapsOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: LogsOpenShift Container Platform: Understanding persistent storage",
					},
					{
						File:    "MavenWrapperDownloader.java",
						Line:    50,
						Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nTwelve-Factor App: Backing servicesOpenShift Container Platform: Input secrets and ConfigMapsTwelve-Factor App: LogsOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: ConfigOpenShift Container Platform: Understanding persistent storage",
					},
					{
						File:    "MavenWrapperDownloader.java",
						Line:    78,
						Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nOpenShift Container Platform: Input secrets and ConfigMapsTwelve-Factor App: ConfigOpenShift Container Platform: Understanding persistent storageTwelve-Factor App: LogsOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: Backing services",
					},
					{
						File:    "MavenWrapperDownloader.java",
						Line:    60,
						Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nTwelve-Factor App: Backing servicesTwelve-Factor App: LogsTwelve-Factor App: ConfigOpenShift Container Platform: Input secrets and ConfigMapsOpenShift Container Platform: Understanding cluster loggingOpenShift Container Platform: Understanding persistent storage",
					},
					{
						File:    "MavenWrapperDownloader.java",
						Line:    55,
						Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nOpenShift Container Platform: Understanding persistent storageOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: LogsTwelve-Factor App: Backing servicesTwelve-Factor App: ConfigOpenShift Container Platform: Input secrets and ConfigMaps",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "configuration-management",
				Rule:        "configuration-management-0200",
				Incidents: []api.Incident{
					{
						File:    "application.properties",
						Line:    0,
						Message: "Application properties file detected",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "technology-usage-javaee",
				Rule:        "javaee-technology-usage-00110",
				Incidents: []api.Incident{
					{
						File:    "AssessmentsResource.java",
						Line:    0,
						Message: "The application uses Bean Validation",
					},
					{
						File:    "TranslatorResource.java",
						Line:    0,
						Message: "The application uses Bean Validation",
					},
					{
						File:    "ApplicationDto.java",
						Line:    0,
						Message: "The application uses Bean Validation",
					},
					{
						File:    "AssessmentCreateCommand.java",
						Line:    0,
						Message: "The application uses Bean Validation",
					},
					{
						File:    "AssessmentSvc.java",
						Line:    0,
						Message: "The application uses Bean Validation",
					},
					{
						File:    "BulkCreateSvc.java",
						Line:    0,
						Message: "The application uses Bean Validation",
					},
					{
						File:    "TranslatorSvc.java",
						Line:    0,
						Message: "The application uses Bean Validation",
					},
				},
			},
			{
				Category:    "information",
				Description: "Info",
				Effort:      0,
				RuleSet:     "DiscoverMavenProjectsRuleProvider",
				Rule:        "DiscoverMavenProjectsRuleProvider_1",
				Incidents: []api.Incident{
					{
						File:    "pom.xml",
						Line:    0,
						Message: "Maven Project Object Model (POM) File",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "CDI"},
		{Name: "HTML"},
		{Name: "Properties"},
		{Name: "JPA entities"},
		{Name: "JAX-RS"},
		{Name: "Bean Validation"},
		{Name: "Application Properties File"},
	},
}
