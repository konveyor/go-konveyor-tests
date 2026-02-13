package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var PetclinicHazelcast = TC{
	Name:        "Petclinic legacy cloud-readiness with tagger and hazelcast custom rules",
	Application: data.PetclinicHazelcast,
	RulesPath:   "/rules",
	CustomRules: []api.RuleSet{
		{
			Rules: []api.Rule{
				{
					File: &api.Ref{
						Name: "/rules/01-hz.windup.yaml",
					},
				},
			},
		},
	},
	Task: Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Analysis: api.Analysis{
		Effort: 20,
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "Embedded Hazelcast",
				Effort:      3,
				RuleSet:     "files",
				Rule:        "hazelcast-cloud-readiness-hz001",
				Incidents: []api.Incident{
					{
						File: "/shared/source/spring-framework-petclinic/src/main/java/org/springframework/samples/petclinic/config/SessionConfiguration.java",
						Line: 77,
						Message: `Consider using Kubernetes specific configuration.
					
							// Example using Kubernetes specific configuration
							
							JoinConfig joinConfig = config.getNetworkConfig().getJoin();
							config.getKubernetesConfig().setEnabled(true)
							 .setProperty("namespace", "namespace")
							 .setProperty("service-name", "hazelcast-service");`,
					},
					{
						File: "/shared/source/spring-framework-petclinic/src/main/java/org/springframework/samples/petclinic/config/SessionConfiguration.java",
						Line: 78,
						Message: `Consider using Kubernetes specific configuration.
					
							// Example using Kubernetes specific configuration
							
							JoinConfig joinConfig = config.getNetworkConfig().getJoin();
							config.getKubernetesConfig().setEnabled(true)
							 .setProperty("namespace", "namespace")
							 .setProperty("service-name", "hazelcast-service");`,
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
						File:    "/shared/source/spring-framework-petclinic/pom.xml",
						Line:    560,
						Message: "The app is trying to access local resource by JDBC, please try to migrate the resource to cloud",
					},
					{
						File:    "/shared/source/spring-framework-petclinic/pom.xml",
						Line:    579,
						Message: "The app is trying to access local resource by JDBC, please try to migrate the resource to cloud",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "org.hamcrest.hamcrest",
				Version:  "2.2",
				Provider: "java",
			},
			{
				Name:     "org.mockito.mockito-junit-jupiter",
				Version:  "4.0.0",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-webmvc",
				Version:  "5.3.12",
				Provider: "java",
			},
			{
				Name:     "ch.qos.logback.logback-classic",
				Version:  "1.2.6",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.validator.hibernate-validator",
				Version:  "6.2.0.Final",
				Provider: "java",
			},
			{
				Name:     "org.webjars.bootstrap",
				Version:  "3.3.6",
				Provider: "java",
			},
			{
				Name:     "org.springframework.data.spring-data-jpa",
				Version:  "2.6.0-RC1",
				Provider: "java",
			},
			{
				Name:     "org.glassfish.jaxb.jaxb-runtime",
				Version:  "2.3.4",
				Provider: "java",
			},
			{
				Name:     "org.apache.taglibs.taglibs-standard-jstlel",
				Version:  "1.2.5",
				Provider: "java",
			},
			{
				Name:     "com.hazelcast.hazelcast",
				Version:  "4.2.2",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-jdbc",
				Version:  "5.3.12",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.tomcat-jasper-el",
				Version:  "9.0.46",
				Provider: "java",
			},
			{
				Name:     "com.h2database.h2",
				Version:  "1.4.200",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-oxm",
				Version:  "5.3.12",
				Provider: "java",
			},
			{
				Name:     "org.aspectj.aspectjrt",
				Version:  "1.9.7",
				Provider: "java",
			},
			{
				Name:     "org.springframework.session.spring-session-core",
				Version:  "2.5.2",
				Provider: "java",
			},
			{
				Name:     "com.jayway.jsonpath.json-path",
				Version:  "2.6.0",
				Provider: "java",
			},
			{
				Name:     "org.slf4j.slf4j-api",
				Version:  "1.7.32",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-core",
				Version:  "2.12.4",
				Provider: "java",
			},
			{
				Name:     "org.mockito.mockito-core",
				Version:  "4.0.0",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-test",
				Version:  "5.3.12",
				Provider: "java",
			},
			{
				Name:     "org.junit.jupiter.junit-jupiter-engine",
				Version:  "5.8.1",
				Provider: "java",
			},
			{
				Name:     "javax.xml.bind.jaxb-api",
				Version:  "2.3.1",
				Provider: "java",
			},
			{
				Name:     "javax.activation.javax.activation-api",
				Version:  "1.2.0",
				Provider: "java",
			},
			{
				Name:     "javax.annotation.javax.annotation-api",
				Version:  "1.3.2",
				Provider: "java",
			},
			{
				Name:     "org.webjars.jquery",
				Version:  "3.5.1",
				Provider: "java",
			},
			{
				Name:     "org.aspectj.aspectjweaver",
				Version:  "1.9.7",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.hibernate-ehcache",
				Version:  "5.6.0.Final",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.tomcat-servlet-api",
				Version:  "9.0.46",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.hibernate-entitymanager",
				Version:  "5.6.0.Final",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.tomcat-jdbc",
				Version:  "9.0.46",
				Provider: "java",
			},
			{
				Name:     "org.assertj.assertj-core",
				Version:  "3.21.0",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-databind",
				Version:  "2.12.4",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-context-support",
				Version:  "5.3.12",
				Provider: "java",
			},
			{
				Name:     "org.webjars.jquery-ui",
				Version:  "1.12.1",
				Provider: "java",
			},
			{
				Name:     "org.springframework.session.spring-session-hazelcast",
				Version:  "2.5.2",
				Provider: "java",
			},
			{
				Name:     "net.bytebuddy.byte-buddy-agent",
				Version:  "1.11.19",
				Provider: "java",
			},
			{
				Name:     "net.minidev.accessors-smart",
				Version:  "2.4.7",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.common.hibernate-commons-annotations",
				Version:  "5.1.2.Final",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.hibernate-core",
				Version:  "5.6.0.Final",
				Provider: "java",
			},
			{
				Name:     "org.apiguardian.apiguardian-api",
				Version:  "1.1.2",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.tomcat-juli",
				Version:  "9.0.46",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.tomcat-el-api",
				Version:  "9.0.46",
				Provider: "java",
			},
			{
				Name:     "org.jboss.jandex",
				Version:  "2.2.3.Final",
				Provider: "java",
			},
			{
				Name:     "org.jboss.logging.jboss-logging",
				Version:  "3.4.2.Final",
				Provider: "java",
			},
			{
				Name:     "org.jboss.spec.javax.transaction.jboss-transaction-api_1.2_spec",
				Version:  "1.1.1.Final",
				Provider: "java",
			},
			{
				Name:     "org.junit.jupiter.junit-jupiter-api",
				Version:  "5.8.1",
				Provider: "java",
			},
			{
				Name:     "org.apache.taglibs.taglibs-standard-spec",
				Version:  "1.2.5",
				Provider: "java",
			},
			{
				Name:     "org.junit.platform.junit-platform-commons",
				Version:  "1.8.1",
				Provider: "java",
			},
			{
				Name:     "org.junit.platform.junit-platform-engine",
				Version:  "1.8.1",
				Provider: "java",
			},
			{
				Name:     "org.apache.taglibs.taglibs-standard-impl",
				Version:  "1.2.5",
				Provider: "java",
			},
			{
				Name:     "net.sf.ehcache.ehcache",
				Version:  "2.10.6",
				Provider: "java",
			},
			{
				Name:     "org.objenesis.objenesis",
				Version:  "3.2",
				Provider: "java",
			},
			{
				Name:     "org.opentest4j.opentest4j",
				Version:  "1.2.0",
				Provider: "java",
			},
			{
				Name:     "org.ow2.asm.asm",
				Version:  "9.1",
				Provider: "java",
			},
			{
				Name:     "net.minidev.json-smart",
				Version:  "2.4.7",
				Provider: "java",
			},
			{
				Name:     "org.springframework.data.spring-data-commons",
				Version:  "2.6.0-RC1",
				Provider: "java",
			},
			{
				Name:     "org.glassfish.jaxb.txw2",
				Version:  "2.3.4",
				Provider: "java",
			},
			{
				Name:     "antlr.antlr",
				Version:  "2.7.7",
				Provider: "java",
			},
			{
				Name:     "net.bytebuddy.byte-buddy",
				Version:  "1.11.20",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-aop",
				Version:  "5.3.11",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-beans",
				Version:  "5.3.11",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-context",
				Version:  "5.3.11",
				Provider: "java",
			},
			{
				Name:     "javax.persistence.javax.persistence-api",
				Version:  "2.2",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-core",
				Version:  "5.3.11",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-expression",
				Version:  "5.3.12",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-jcl",
				Version:  "5.3.9",
				Provider: "java",
			},
			{
				Name:     "jakarta.xml.bind.jakarta.xml.bind-api",
				Version:  "2.3.3",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-orm",
				Version:  "5.3.11",
				Provider: "java",
			},
			{
				Name:     "jakarta.validation.jakarta.validation-api",
				Version:  "2.0.2",
				Provider: "java",
			},
			{
				Name:     "com.sun.istack.istack-commons-runtime",
				Version:  "3.0.12",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-tx",
				Version:  "5.3.11",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-web",
				Version:  "5.3.12",
				Provider: "java",
			},
			{
				Name:     "com.sun.activation.jakarta.activation",
				Version:  "1.2.2",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-annotations",
				Version:  "2.12.4",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.classmate",
				Version:  "1.5.1",
				Provider: "java",
			},
			{
				Name:     "ch.qos.logback.logback-core",
				Version:  "1.2.6",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Other"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Validation"}},
		{Name: "JSP Page", Category: api.Ref{Name: "Web"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSP Page", Category: api.Ref{Name: "Java EE"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "JSP Page", Category: api.Ref{Name: "View"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Bean Validation", Category: api.Ref{Name: "Store"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Spring JMX", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring JMX", Category: api.Ref{Name: "Observability"}},
		{Name: "Spring JMX", Category: api.Ref{Name: "Sustain"}},

	},
}
