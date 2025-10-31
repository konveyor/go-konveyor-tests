package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
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
		Effort: 49,
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "The package 'javax' has been replaced by 'jakarta'.",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-import-00001",
				Incidents: []api.Incident{
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/AirportCodeMapping.java",
					// 	Line:    4,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/AirportCodeMapping.java",
					// 	Line:    5,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Booking.java",
					// 	Line:    5,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Booking.java",
					// 	Line:    6,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Booking.java",
					// 	Line:    7,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Booking.java",
					// 	Line:    8,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/BookingPK.java",
					// 	Line:    4,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/BookingPK.java",
					// 	Line:    5,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Customer.java",
					// 	Line:    4,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Customer.java",
					// 	Line:    5,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Customer.java",
					// 	Line:    6,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Customer.java",
					// 	Line:    7,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/CustomerAddress.java",
					// 	Line:    4,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/CustomerSession.java",
					// 	Line:    5,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/CustomerSession.java",
					// 	Line:    6,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/CustomerSession.java",
					// 	Line:    7,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Flight.java",
					// 	Line:    6,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/Flight.java",
					// 	Line:    7,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/FlightPK.java",
					// 	Line:    4,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/FlightPK.java",
					// 	Line:    5,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/FlightSegment.java",
					// 	Line:    4,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/entities/FlightSegment.java",
					// 	Line:    5,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/BookingServiceImpl.java",
					// 	Line:    14,
					// 	Message: "Replace the `javax.annotation` import statement with `jakarta.annotation`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/BookingServiceImpl.java",
					// 	Line:    15,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/BookingServiceImpl.java",
					// 	Line:    16,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/CustomerServiceImpl.java",
					// 	Line:    10,
					// 	Message: "Replace the `javax.annotation` import statement with `jakarta.annotation`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/CustomerServiceImpl.java",
					// 	Line:    11,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/FlightServiceImpl.java",
					// 	Line:    14,
					// 	Message: "Replace the `javax.annotation` import statement with `jakarta.annotation`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/FlightServiceImpl.java",
					// 	Line:    15,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
					// {
					// 	File:    "/shared/bin/java-project/src/main/java/com/acmeair/jpa/service/FlightServiceImpl.java",
					// 	Line:    16,
					// 	Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					// },
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
		},
		Dependencies: []api.TechDependency{
			{
				Name: "org.springframework.spring-beans-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-asm-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "net.sf.cglib-2.2.2",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: ".aspectjweaver-1.6.8.jar",
				Version: "",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-web-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-core-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-tx-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: ".asm-3.3.1.jar",
				Version: "",
				Provider: "java",
			},
			{
				Name: ".aopalliance-1.0.jar",
				Version: "",
				Provider: "java",
			},
			{
				Name: "net.wasdev.wlp.sample.acmeair-services-jpa",
				Version: "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name: "net.wasdev.wlp.sample.acmeair-services",
				Version: "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-context-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-expression-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "org.aspectj.aspectjrt-1.6.8",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-aop-3.1.2.RELEASE",
				Version: "Unknown",
				Provider: "java",
			},
			{
				Name: "net.wasdev.wlp.sample.acmeair-common",
				Version: "1.0-SNAPSHOT",
				Provider: "java",
			},
			{
				Name: "commons-logging.commons-logging",
				Version: "1.1.1",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "RMI", Category: api.Ref{Name: "Other"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "RMI", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "RMI", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Java EE"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Other"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Connect"}},
	},
}
