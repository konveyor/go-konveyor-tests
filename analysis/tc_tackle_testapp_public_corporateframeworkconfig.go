package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPublicCorpFrameworkConfig = TC{
	Name:        "Tackle Testapp public corporate-framework-config",
	Application: TackleTestApp,
	Task:        Analyze,
	WithDeps:    true,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=linux",
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/source=java",
		},
	},
	Analysis: api.Analysis{
		Effort: 2,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "Hardcoded IP Address\nWhen migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
				Effort: 1,
				RuleSet: "discovery-rules",
				Rule: "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/src/main/resources/persistence.properties",
						Line: 2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
					{
						File: "/addon/source/tackle-testapp-public/target/classes/persistence.properties",
						Line: 2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
				},
			},
			{
				Category: "potential",
				Description: "JBoss API reference",
				Effort: 0,
				RuleSet: "eap6/resin",
				Rule: "generic-catchall-00700",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/controller/CustomerController.java",
						Line: 3,
						Message: "`org.jboss..logging.Logger;import io` reference found. No specific details available.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/exception/handler/ExceptionHandlingController.java",
						Line: 4,
						Message: "`org.jboss..logging.Logger;import io` reference found. No specific details available.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/service/CustomerService.java",
						Line: 4,
						Message: "`org.jboss..logging.Logger;import io` reference found. No specific details available.",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name: "com.fasterxml.jackson.jackson-bom",
				Version: "${jackson.version}",
				Provider: "",
			},
			{
				Name: "org.springframework.data.spring-data-bom",
				Version: "${spring-data.version}",
				Provider: "",
			},
			{
				Name: "org.apache.tomcat.tomcat-servlet-api",
				Version: "${tomcat.version}",
				Provider: "",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-core",
				Version: "",
				Provider: "",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-databind",
				Version: "",
				Provider: "",
			},
			{
				Name: "org.springframework.data.spring-data-jpa",
				Version: "",
				Provider: "",
			},
			{
				Name: "org.springframework.spring-jdbc",
				Version: "${spring-framework.version}",
				Provider: "",
			},
			{
				Name: "org.springframework.spring-webmvc",
				Version: "${spring-framework.version}",
				Provider: "",
			},
			{
				Name: "org.springframework.spring-web",
				Version: "${spring-framework.version}",
				Provider: "",
			},
			{
				Name: "org.springframework.boot.spring-boot-starter-actuator",
				Version: "2.5.0",
				Provider: "",
			},
			{
				Name: "org.apache.tomcat.tomcat-jdbc",
				Version: "${tomcat.version}",
				Provider: "",
			},
			{
				Name: "org.hibernate.hibernate-entitymanager",
				Version: "${hibernate.version}",
				Provider: "",
			},
			{
				Name: "org.hibernate.validator.hibernate-validator",
				Version: "${hibernate-validator.version}",
				Provider: "",
			},
			{
				Name: "ch.qos.logback.logback-classic",
				Version: "1.1.7",
				Provider: "",
			},
			{
				Name: "com.oracle.database.jdbc.ojdbc11",
				Version: "21.1.0.0",
				Provider: "",
			},
			{
				Name: "org.postgresql.postgresql",
				Version: "42.2.23",
				Provider: "",
			},
			{
				Name: "com.h2database.h2",
				Version: "2.1.214",
				Provider: "",
			},
			{
				Name: "io.konveyor.demo.configuration-utils",
				Version: "1.0.0",
				Provider: "",
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
