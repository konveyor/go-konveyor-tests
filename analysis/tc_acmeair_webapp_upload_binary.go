package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/shared/api"
)

var AcmeairWebappBinary = TC{
	Name:        "acmeair-webapp",
	Application: data.UploadBinary,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=jakarta-ee",
		},
	},
	Binary:   true,
	Artifact: "/binary/acmeair-webapp-1.0-SNAPSHOT.war",
	Analysis: api.Analysis{
		Effort: 73,
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "The package 'javax' has been replaced by 'jakarta'.",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-import-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    8,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    9,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    10,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    11,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    12,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    13,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    14,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    15,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/BookingsREST.java",
						Line:    16,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    6,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    7,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    8,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    9,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    10,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    11,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    12,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    13,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    14,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/CustomerREST.java",
						Line:    15,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/FlightsREST.java",
						Line:    7,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/FlightsREST.java",
						Line:    8,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/FlightsREST.java",
						Line:    9,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/FlightsREST.java",
						Line:    10,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/FlightsREST.java",
						Line:    11,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoaderREST.java",
						Line:    19,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoaderREST.java",
						Line:    20,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoaderREST.java",
						Line:    21,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    5,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    6,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    7,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    8,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    9,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    10,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    11,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    12,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    13,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    14,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/LoginREST.java",
						Line:    15,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    6,
						Message: "Replace the `javax.annotation` import statement with `jakarta.annotation`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    7,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    8,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    9,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    10,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    11,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    12,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    13,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    14,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
					{
						File:    "/shared/bin/java-project/src/main/java/com/acmeair/web/RESTCookieSessionFilter.java",
						Line:    15,
						Message: "Replace the `javax.servlet` import statement with `jakarta.servlet`",
					},
				},
			},
			{
				Category:    "potential",
				Description: "web.xml element references a javax-prefixed class name",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-servlet-00130",
				Incidents: []api.Incident{
					{
						File:    "/shared/bin/java-project/src/main/webapp/WEB-INF/web.xml",
						Line:    47,
						Message: "web.xml element references a javax-prefixed class name",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Version of Spring not compatible with Jakarta EE 9+",
				Effort:      3,
				RuleSet:     "jakarta-ee9",
				Rule:        "spring-components-00002",
				Incidents: []api.Incident{
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    74,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    80,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    86,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    92,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    98,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    104,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    110,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
					{
						File:    "/shared/bin/java-project/pom.xml",
						Line:    116,
						Message: "Version 6.0.0 is the minimum version of Spring that is Jakarta EE 9+ compatible",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "org.springframework.spring-beans",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-asm",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.aspectj.aspectjrt",
				Version:  "1.6.8",
				Provider: "java",
			},
			{
				Name:     "net.wasdev.wlp.sample.acmeair-services-jpa",
				Version:  "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-web",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-core",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-tx",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "net.wasdev.wlp.sample.acmeair-common",
				Version:  "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     "commons-logging.commons-logging",
				Version:  "1.1.1",
				Provider: "java",
			},
			{
				Name:     "cglib.cglib",
				Version:  "2.2.2",
				Provider: "java",
			},
			{
				Name:     "asm.asm",
				Version:  "3.3.1",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-context",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-expression",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "net.wasdev.wlp.sample.acmeair-services",
				Version:  "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-aop",
				Version:  "3.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "aopalliance.aopalliance",
				Version:  "1.0",
				Provider: "java",
			},
			{
				Name:     "org.aspectj.aspectjweaver",
				Version:  "1.6.8",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Other"}},
		{Name: "RMI", Category: api.Ref{Name: "Other"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Web"}},
		{Name: "Spring Web", Category: api.Ref{Name: "View"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Execute"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Embedded"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Java EE"}},
		{Name: "RMI", Category: api.Ref{Name: "Connect"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "RMI", Category: api.Ref{Name: "Java EE"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Connect"}},
	},
}
