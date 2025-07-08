package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var DaytraderWithDeps = TC{
	Name:        "Daytrader",
	Application: data.Daytrader,
	WithDeps:    true,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=quarkus",
		},
	},
	Analysis: api.Analysis{
		Effort: 318,
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "File system - java.net.URL/URI",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00002",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/PingReentryServlet.java",
						Line: 91,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/PingServlet2PDF.java",
						Line: 86,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Local HTTP Calls",
				Effort:      7,
				RuleSet:     "cloud-readiness",
				Rule:        "localhost-http-00001",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/PingServlet2PDF.java",
						Line: 85,
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
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/TradeBuildDB.java",
						Line: 43,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/TradeScenarioServlet.java",
						Line: 125,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Adopt Maven Surefire plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00040",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7/pom.xml",
						Line: 2,
					},
					{
						File: "/shared/source/sample/pom.xml",
						Line: 5,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Add Maven profile to run the Quarkus native build",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00060",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7/pom.xml",
						Line: 2,
					},
					{
						File: "/shared/source/sample/pom.xml",
						Line: 5,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JMS' Queue must be replaced with an Emitter",
				Effort:      3,
				RuleSet:     "quarkus/springboot",
				Rule:        "jms-to-reactive-quarkus-00030",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 2044,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 2047,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 2095,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 33,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 37,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 79,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBQueue.java",
						Line: 24,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBQueue.java",
						Line: 58,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Remote EJBs are not supported in Quarkus",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "remote-ejb-to-quarkus-00000",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBRemote.java",
						Line: 26,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "@Stateful annotation must be replaced",
				Effort:      3,
				RuleSet:     "quarkus/springboot",
				Rule:        "ee-to-quarkus-00010",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/PingEJBLocal.java",
						Line: 24,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Method should be marked as @Transactional",
				Effort:      3,
				RuleSet:     "quarkus/springboot",
				Rule:        "ee-to-quarkus-00020",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 1421,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "The expected project artifact's extension is `jar`",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00000",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-web/pom.xml",
						Line: 6,
					},
					{
						File: "/shared/source/sample/daytrader-ee7/pom.xml",
						Line: 29,
					},
					{
						File: "/shared/source/sample/pom.xml",
						Line: 9,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Adopt Quarkus Maven plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00020",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7/pom.xml",
						Line: 2,
					},
					{
						File: "/shared/source/sample/pom.xml",
						Line: 5,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "@MessageDriven - EJBs are not supported in Quarkus",
				Effort:      3,
				RuleSet:     "quarkus/springboot",
				Rule:        "jms-to-reactive-quarkus-00010",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTBroker3MDB.java",
						Line: 39,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTStreamer3MDB.java",
						Line: 40,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Adopt Maven Compiler plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00030",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7/pom.xml",
						Line: 2,
					},
					{
						File: "/shared/source/sample/pom.xml",
						Line: 5,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Adopt Maven Failsafe plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00050",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7/pom.xml",
						Line: 2,
					},
					{
						File: "/shared/source/sample/pom.xml",
						Line: 5,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JMS is not supported in Quarkus",
				Effort:      5,
				RuleSet:     "quarkus/springboot",
				Rule:        "jms-to-reactive-quarkus-00050",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 30,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 31,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 32,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 33,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 34,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 35,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTBroker3MDB.java",
						Line: 27,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTBroker3MDB.java",
						Line: 28,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTBroker3MDB.java",
						Line: 29,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTStreamer3MDB.java",
						Line: 28,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTStreamer3MDB.java",
						Line: 29,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/DTStreamer3MDB.java",
						Line: 30,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 41,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 36,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 37,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 38,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 39,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 40,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBQueue.java",
						Line: 21,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBQueue.java",
						Line: 22,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBQueue.java",
						Line: 23,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBQueue.java",
						Line: 24,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBQueue.java",
						Line: 25,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBTopic.java",
						Line: 21,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBTopic.java",
						Line: 22,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBTopic.java",
						Line: 23,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBTopic.java",
						Line: 24,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBTopic.java",
						Line: 25,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/websocket/MarketSummaryWebSocket.java",
						Line: 25,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/websocket/RecentStockChangeList.java",
						Line: 25,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Adopt Quarkus BOM",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00010",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/pom.xml",
						Line: 3,
					},
					{
						File: "/shared/source/sample/daytrader-ee7/pom.xml",
						Line: 2,
					},
					{
						File: "/shared/source/sample/pom.xml",
						Line: 5,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JMS' Topic must be replaced with an Emitter",
				Effort:      3,
				RuleSet:     "quarkus/springboot",
				Rule:        "jms-to-reactive-quarkus-00040",
				Incidents: []api.Incident{
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 2062,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 2065,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 2099,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/direct/TradeDirect.java",
						Line: 35,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 40,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-ejb/src/main/java/com/ibm/websphere/samples/daytrader/ejb3/TradeSLSBBean.java",
						Line: 76,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBTopic.java",
						Line: 25,
					},
					{
						File: "/shared/source/sample/daytrader-ee7-web/src/main/java/com/ibm/websphere/samples/daytrader/web/prims/ejb3/PingServlet2MDBTopic.java",
						Line: 58,
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "taglibs.standard",
				Version:  "1.1.1",
				Provider: "java",
			},
			{
				Name:     "org.apache.derby.derby",
				Version:  "10.14.2.0",
				Provider: "java",
			},
			{
				Name:     "net.wasdev.wlp.sample.daytrader-ee7-web",
				Version:  "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     "net.wasdev.wlp.sample.daytrader-ee7-ejb",
				Version:  "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     "javax.javaee-api",
				Version:  "7.0",
				Provider: "java",
			},
			{
				Name:     "javax.annotation.javax.annotation-api",
				Version:  "1.3.2",
				Provider: "java",
			},
			{
				Name:     "com.sun.mail.javax.mail",
				Version:  "1.5.0",
				Provider: "java",
			},
			{
				Name:     "javax.activation.activation",
				Version:  "1.1",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "CDI", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "JSF", Category: api.Ref{Name: "MVC"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Other"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "EJB Timer", Category: api.Ref{Name: "Processing"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Processing"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Validation"}},
		{Name: "JSF Page", Category: api.Ref{Name: "Web"}},
		{Name: "JSP Page", Category: api.Ref{Name: "Web"}},
		{Name: "WebSocket", Category: api.Ref{Name: "Web"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "CDI", Category: api.Ref{Name: "Java EE"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Connect"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Store"}},
		{Name: "JSF", Category: api.Ref{Name: "View"}},
		{Name: "JSF XML", Category: api.Ref{Name: "View"}},
		{Name: "EJB Timer", Category: api.Ref{Name: "Execute"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSP Page", Category: api.Ref{Name: "View"}},
		{Name: "EJB Timer", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSP Page", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "CDI XML", Category: api.Ref{Name: "Execute"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "WebSocket", Category: api.Ref{Name: "View"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Java EE"}},
		{Name: "CDI XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF", Category: api.Ref{Name: "Embedded"}},
		{Name: "CDI", Category: api.Ref{Name: "Execute"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Execute"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "JSF Page", Category: api.Ref{Name: "View"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Web"}},
		{Name: "JSF Page", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "CDI XML", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "WebSocket", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
	},
}
