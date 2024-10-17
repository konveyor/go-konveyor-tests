package analysis

import "github.com/konveyor/tackle2-hub/api"

var Daytrader = TC{
	Name: "Daytrader",
	Application: api.Application{
		Name: "Daytrader 7 EE application",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/WASdev/sample.daytrader7.git",
		},
	},
	WithDeps: true,
	Task:     Analyze,
	ReportContent: map[string][]string{
		"/windup/report/index.html": {
			"0\nstory points",
			"6\nInformation",
		},
	},
	Targets: []string{
		"konveyor.io/target=cloud-readiness",
		"konveyor.io/target=quarkus",
	},
	Analysis: api.Analysis{
		Effort: 10,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Effort: 1,
				Description: "File system - Java IO",
				RuleSet: "cloud-readiness",
				Rule: "local-storage-00001",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/TradeScenarioServlet.java",
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
				},
			},
			{
				Category: "mandatory",
				Effort: 1,
				Description: "File system - java.net.URL/URI",
				RuleSet: "cloud-readiness",
				Rule: "local-storage-00002",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/PingReentryServlet.java",
						Line: 91,
						Message: `An application running inside a container could lose access to a file in local storage.

 Recommendations

 The following recommendations depend on the function of the file in local storage:

 * Logging: Log to standard output and use a centralized log collector to analyze the logs.
 * Caching: Use a cache backing service.
 * Configuration: Store configuration settings in environment variables so that they can be updated without code changes.
 * Data storage: Use a database backing service for relational data or use a persistent data storage system.
 * Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.`,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/PingServlet2PDF.java",
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
				},
			},
			{
				Category: "mandatory",
				Effort: 7,
				Description: "Local HTTP Calls",
				RuleSet: "cloud-readiness",
				Rule: "localhost-http-00001",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/PingServlet2PDF.java",
						Line: 85,
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "CDI", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "CDI XML", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "EJB Timer",Category: api.Ref{Name: "Processing"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Processing"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Validation"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Store"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "CDI", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB Timer", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "CDI XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "CDI", Category: api.Ref{Name: "Execute"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Execute"}},
		{Name: "EJB Timer", Category: api.Ref{Name: "Execute"}},
		{Name: "CDI XML", Category: api.Ref{Name: "Execute"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Web"}},
		{Name: "JSF XML", Category: api.Ref{Name: "View"}},
	},
}
