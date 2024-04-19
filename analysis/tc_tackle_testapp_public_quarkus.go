package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPublicQuarkus = TC{
	Name:        "Tackle Testapp with Quarkus target",
	Application: TackleTestApp,
	Task:        Analyze,
	WithDeps:    true,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=quarkus",
		},
	},
	Identities: []api.Identity{
		MavenPublic, // Tackle Testapp public Maven registry expects MAVEN_TESTAPP_USER and MAVEN_TESTAPP_TOKEN env variables.

	},
	Analysis: api.Analysis{
		Effort: 39,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "Replace the Spring Web artifact with Quarkus 'spring-web' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "springboot-web-to-quarkus-00000",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the Spring Web artifact with Quarkus `spring-web` extension",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Adopt Maven Surefire plugin",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "javaee-pom-to-quarkus-00040",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/configuration-utils/pom.xml",
						Line: 6,
						Message: "Use the Maven Surefire plugin adding the following sections to the `pom.xml` file",
					},
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 4,
						Message: "Use the Maven Surefire plugin adding the following sections to the `pom.xml` file:",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Adopt Maven Failsafe plugin",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "javaee-pom-to-quarkus-00050",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/configuration-utils/pom.xml",
						Line: 6,
						Message: "Use the Maven Failsafe plugin adding the following sections to the `pom.xml` file:",
					},
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 4,
						Message: "Use the Maven Failsafe plugin adding the following sections to the `pom.xml` file:",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Configure Quarkus hibernate-orm",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "javaee-pom-to-quarkus-00070",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 99,
						Message: "Configure Quarkus 'hibernate-orm` and jakarta persistence.",
					},
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 104,
						Message: "Configure Quarkus 'hibernate-orm` and jakarta persistence.",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the 'ojdbc11' dependency with Quarkus 'quarkus-jdbc-oracle' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "quarkus-jdbc-oracle-00000",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the `com.oracle.database.jdbc:ojdbc11` dependency with the Quarkus dependency `io.quarkus:quarkus-jdbc-oracle`",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the Spring Boot Actuator dependency with Quarkus Smallrye Health extension",
				Effort: 5,
				RuleSet: "quarkus/springboot",
				Rule: "springboot-actuator-to-quarkus-0100",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the Spring Boot Actuator dependency with Quarkus Smallrye Health extension.",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the SpringBoot Data JPA artifact with Quarkus 'spring-data-jpa' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "springboot-jpa-to-quarkus-00000",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the SpringBoot JPA artifact with Quarkus `spring-data-jpa` extension",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Adopt Quarkus BOM",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "javaee-pom-to-quarkus-00010",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/configuration-utils/pom.xml",
						Line: 6,
						Message: "Use the Quarkus BOM to omit the version of the different Quarkus dependencies.",
					},
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 4,
						Message: "Use the Quarkus BOM to omit the version of the different Quarkus dependencies.",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Adopt Maven Compiler plugin",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "javaee-pom-to-quarkus-00030",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/configuration-utils/pom.xml",
						Line: 6,
						Message: "Use the Maven Compiler plugin adding the following sections to the `pom.xml` file:",
					},
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 4,
						Message: "Use the Maven Compiler plugin adding the following sections to the `pom.xml` file:",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the 'hibernate-validator' dependency with Quarkus 'quarkus-hibernate-validator' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "quarkus-hibernate-validator-00000",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the `org.hibernate.validator:hibernate-validator` dependency with the Quarkus dependency `io.quarkus:quarkus-hibernate-validator`",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the Micrometer dependency with Quarkus Microprofile 'metrics' extension",
				Effort: 5,
				RuleSet: "quarkus/springboot",
				Rule: "springboot-metrics-to-quarkus-0100",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the Micrometer dependency with Quarkus Microprofile 'metrics' extension.",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the Micrometer code with Microprofile Metrics code",
				Effort: 5,
				RuleSet: "quarkus/springboot",
				Rule: "springboot-metrics-to-quarkus-0200",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the Micrometer code with Microprofile Metrics code",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Adopt Quarkus Maven plugin",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "javaee-pom-to-quarkus-00020",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/configuration-utils/pom.xml",
						Line: 6,
						Message: "Use the Quarkus Maven plugin adding the following sections to the `pom.xml` file:",
					},
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 4,
						Message: "Use the Quarkus Maven plugin adding the following sections to the `pom.xml` file:",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Add Maven profile to run the Quarkus native build",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "javaee-pom-to-quarkus-00060",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/configuration-utils/pom.xml",
						Line: 6,
						Message: "Leverage a Maven profile to run the Quarkus native build adding the following section to the `pom.xml` file:",
					},
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 4,
						Message: "Leverage a Maven profile to run the Quarkus native build adding the following section to the `pom.xml` file:",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the 'postgresql' dependency with Quarkus 'quarkus-jdbc-postgresql' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "quarkus-jdbc-postgresql-00010",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the `org.postgresql:postgresql` dependency with the Quarkus dependency `io.quarkus:quarkus-jdbc-postgresql`",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the 'hibernate-validator' dependency with Quarkus 'quarkus-hibernate-validator' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "quarkus-hibernate-validator-00010",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the `org.hibernate.validator:hibernate-validator` dependency with the Quarkus dependency `io.quarkus:quarkus-hibernate-validator`",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the 'ojdbc11' dependency with Quarkus 'quarkus-jdbc-oracle' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "quarkus-jdbc-oracle-00010",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the `com.oracle.database.jdbc:ojdbc11` dependency with the Quarkus dependency `io.quarkus:quarkus-jdbc-oracle`",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the 'postgresql' dependency with Quarkus 'quarkus-jdbc-postgresql' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "quarkus-jdbc-postgresql-00000",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the `org.postgresql:postgresql` dependency with the Quarkus dependency `io.quarkus:quarkus-jdbc-postgresql`",
					},
				},
			},
			{
				Category: "mandatory",
				Description: "Replace the SpringBoot artifact with Quarkus 'spring-boot-properties' extension",
				Effort: 1,
				RuleSet: "quarkus/springboot",
				Rule: "springboot-properties-to-quarkus-00000",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/pom.xml",
						Line: 0,
						Message: "Replace the SpringBoot artifact with Quarkus `spring-boot-properties` extension",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name: "org.postgresql.postgresql",
				Version: "42.2.23",
				Provider: "java",
			},
			{
				Name: "org.springframework.boot.spring-boot-starter-actuator",
				Version: "2.5.0",
				Provider: "java",
			},
			{
				Name: "com.oracle.database.jdbc.ojdbc11",
				Version: "21.1.0.0",
				Provider: "java",
			},
			{
				Name: "org.apache.tomcat.tomcat-jdbc",
				Version: "9.0.46",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-web",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.hibernate.validator.hibernate-validator",
				Version: "6.2.0.Final",
				Provider: "java",
			},
			{
				Name: "org.apache.tomcat.tomcat-servlet-api",
				Version: "9.0.46",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-jdbc",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.data.spring-data-jpa",
				Version: "2.5.1",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-webmvc",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "ch.qos.logback.logback-classic",
				Version: "1.1.7",
				Provider: "java",
			},
			{
				Name: "io.konveyor.demo.configuration-utils",
				Version: "1.0.0",
				Provider: "java",
			},
			{
				Name: "org.hibernate.hibernate-entitymanager",
				Version: "5.4.32.Final",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-databind",
				Version: "2.12.3",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-core",
				Version: "2.12.3",
				Provider: "java",
			},
			{
				Name: "com.h2database.h2",
				Version: "2.1.214",
				Provider: "java",
			},
			{
				Name: "org.hibernate.common.hibernate-commons-annotations",
				Version: "5.1.2.Final",
				Provider: "java",
			},
			{
				Name: "com.sun.xml.fastinfoset.FastInfoset",
				Version: "1.2.15",
				Provider: "java",
			},
			{
				Name: "javax.xml.bind.jaxb-api",
				Version: "2.3.1",
				Provider: "java",
			},
			{
				Name: "net.bytebuddy.byte-buddy",
				Version: "1.10.22",
				Provider: "java",
			},
			{
				Name: "org.apache.logging.log4j.log4j-api",
				Version: "2.14.1",
				Provider: "java",
			},
			{
				Name: "org.apache.logging.log4j.log4j-to-slf4j",
				Version: "2.14.1",
				Provider: "java",
			},
			{
				Name: "javax.activation.javax.activation-api",
				Version: "1.2.0",
				Provider: "java",
			},
			{
				Name: "org.apache.tomcat.tomcat-juli",
				Version: "9.0.46",
				Provider: "java",
			},
			{
				Name: "jakarta.validation.jakarta.validation-api",
				Version: "2.0.2",
				Provider: "java",
			},
			{
				Name: "org.aspectj.aspectjrt",
				Version: "1.9.6",
				Provider: "java",
			},
			{
				Name: "org.checkerframework.checker-qual",
				Version: "3.5.0",
				Provider: "java",
			},
			{
				Name: "org.dom4j.dom4j",
				Version: "2.1.3",
				Provider: "java",
			},
			{
				Name: "org.glassfish.jaxb.jaxb-runtime",
				Version: "2.3.1",
				Provider: "java",
			},
			{
				Name: "org.glassfish.jaxb.txw2",
				Version: "2.3.1",
				Provider: "java",
			},
			{
				Name: "org.hdrhistogram.HdrHistogram",
				Version: "2.1.12",
				Provider: "java",
			},
			{
				Name: "antlr.antlr",
				Version: "2.7.7",
				Provider: "java",
			},
			{
				Name: "org.hibernate.hibernate-core",
				Version: "5.4.32.Final",
				Provider: "java",
			},
			{
				Name: "jakarta.annotation.jakarta.annotation-api",
				Version: "1.3.5",
				Provider: "java",
			},
			{
				Name: "io.micrometer.micrometer-core",
				Version: "1.7.0",
				Provider: "java",
			},
			{
				Name: "org.javassist.javassist",
				Version: "3.27.0-GA",
				Provider: "java",
			},
			{
				Name: "org.jboss.jandex",
				Version: "2.2.3.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.logging.jboss-logging",
				Version: "3.4.1.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.spec.javax.transaction.jboss-transaction-api_1.2_spec",
				Version: "1.1.1.Final",
				Provider: "java",
			},
			{
				Name: "org.jvnet.staxex.stax-ex",
				Version: "1.8",
				Provider: "java",
			},
			{
				Name: "org.latencyutils.LatencyUtils",
				Version: "2.0.3",
				Provider: "java",
			},
			{
				Name: "javax.persistence.javax.persistence-api",
				Version: "2.2",
				Provider: "java",
			},
			{
				Name: "org.slf4j.jul-to-slf4j",
				Version: "1.7.30",
				Provider: "java",
			},
			{
				Name: "org.slf4j.slf4j-api",
				Version: "1.7.26",
				Provider: "java",
			},
			{
				Name: "org.springframework.boot.spring-boot",
				Version: "2.5.0",
				Provider: "java",
			},
			{
				Name: "org.springframework.boot.spring-boot-actuator",
				Version: "2.5.0",
				Provider: "java",
			},
			{
				Name: "org.springframework.boot.spring-boot-actuator-autoconfigure",
				Version: "2.5.0",
				Provider: "java",
			},
			{
				Name: "org.springframework.boot.spring-boot-autoconfigure",
				Version: "2.5.0",
				Provider: "java",
			},
			{
				Name: "org.springframework.boot.spring-boot-starter",
				Version: "2.5.0",
				Provider: "java",
			},
			{
				Name: "com.sun.istack.istack-commons-runtime",
				Version: "3.0.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.boot.spring-boot-starter-logging",
				Version: "2.5.0",
				Provider: "java",
			},
			{
				Name: "org.springframework.data.spring-data-commons",
				Version: "2.5.1",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.datatype.jackson-datatype-jsr310",
				Version: "2.12.3",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-aop",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-beans",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-context",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-core",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-expression",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-jcl",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-annotations",
				Version: "2.12.3",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-orm",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "org.springframework.spring-tx",
				Version: "5.3.7",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.classmate",
				Version: "1.5.1",
				Provider: "java",
			},
			{
				Name: "ch.qos.logback.logback-core",
				Version: "1.1.7",
				Provider: "java",
			},
			{
				Name: "org.yaml.snakeyaml",
				Version: "1.28",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Micrometer", Category: api.Ref{Name: "Integration"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "Spring Data JPA", Category: api.Ref{Name: "Persistence"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Web"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Micrometer", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Data JPA", Category: api.Ref{Name: "Store"}},
		{Name: "Spring Web", Category: api.Ref{Name: "View"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Web", Category: api.Ref{Name: "Embedded"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Spring Data JPA", Category: api.Ref{Name: "Embedded"}},
		{Name: "Micrometer", Category: api.Ref{Name: "Execute"}},
		{Name: "Spring DI", Category: api.Ref{Name: "Execute"}},
	},
}
