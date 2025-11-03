package analysis

import "github.com/konveyor/tackle2-hub/api"

var BookServerSource = TC{
	SkipTest: SkipTestConfig{
		Reason: "Skip test. Fails due to https://github.com/konveyor/analyzer-lsp/issues/936",
		Skip:   true,
	},
	Name: "book-server_source",
	Task: Analyze,
	Analysis: api.Analysis{
		Dependencies: []api.TechDependency{
			{
				Name:     "com.fasterxml.jackson.module.jackson-module-jaxb-annotations",
				Version:  "2.9.5",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.dataformat.jackson-dataformat-xml",
				Version:  "2.9.5",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-databind",
				Version:  "2.9.5",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.woodstox.woodstox-core",
				Version:  "5.0.3",
				Provider: "java",
			},
			{
				Name:     "org.codehaus.woodstox.stax2-api",
				Version:  "3.1.4",
				Provider: "java",
			},
			{
				Name:     "org.projectlombok.lombok",
				Version:  "1.18.2",
				Provider: "java",
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter-web",
				Version:  "2.1.0.RELEASE",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.datatype.jackson-datatype-jsr310",
				Version:  "2.9.5",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.validator.hibernate-validator",
				Version:  "6.0.13.Final",
				Provider: "java",
			},
			{
				Name:     "org.springframework.boot.spring-boot",
				Version:  "2.1.0.RELEASE",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.module.jackson-module-parameter-names",
				Version:  "2.9.7",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-core",
				Version:  "2.9.5",
				Provider: "java",
			},
			{
				Name:     "javax.annotation.javax.annotation-api",
				Version:  "1.3.2",
				Provider: "java",
			},
			{
				Name:     "javax.validation.validation-api",
				Version:  "2.0.1.Final",
				Provider: "java",
			},
			{
				Name:     "org.apache.logging.log4j.log4j-api",
				Version:  "2.11.1",
				Provider: "java",
			},
			{
				Name:     "org.apache.logging.log4j.log4j-to-slf4j",
				Version:  "2.11.1",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.embed.tomcat-embed-core",
				Version:  "9.0.12",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.embed.tomcat-embed-el",
				Version:  "9.0.12",
				Provider: "java",
			},
			{
				Name:     "org.apache.tomcat.embed.tomcat-embed-websocket",
				Version:  "9.0.12",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-annotations",
				Version:  "2.9.0",
				Provider: "java",
			},
			{
				Name:     "ch.qos.logback.logback-classic",
				Version:  "1.2.3",
				Provider: "java",
			},
			{
				Name:     "org.jboss.logging.jboss-logging",
				Version:  "3.3.2.Final",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.classmate",
				Version:  "1.3.4",
				Provider: "java",
			},
			{
				Name:     "org.slf4j.jul-to-slf4j",
				Version:  "1.7.25",
				Provider: "java",
			},
			{
				Name:     "org.slf4j.slf4j-api",
				Version:  "1.7.25",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.datatype.jackson-datatype-jdk8",
				Version:  "2.9.7",
				Provider: "java",
			},
			{
				Name:     "org.springframework.boot.spring-boot-autoconfigure",
				Version:  "2.1.0.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter",
				Version:  "2.1.0.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter-json",
				Version:  "2.1.0.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter-logging",
				Version:  "2.1.0.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.boot.spring-boot-starter-tomcat",
				Version:  "2.1.0.RELEASE",
				Provider: "java",
			},
			{
				Name:     "ch.qos.logback.logback-core",
				Version:  "1.2.3",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-aop",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-beans",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-context",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-core",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-expression",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-jcl",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-web",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-webmvc",
				Version:  "5.1.2.RELEASE",
				Provider: "java",
			},
			{
				Name:     "org.yaml.snakeyaml",
				Version:  "1.23",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Spring Boot Configuration", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Boot Component Scan", Category: api.Ref{Name: "Configuration Management"}},
		{Name: "Spring Boot Configuration", Category: api.Ref{Name: "Configuration Management"}},
		{Name: "Spring Boot Auto-configuration", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Boot Auto-configuration", Category: api.Ref{Name: "Sustain"}},
		{Name: "Spring Boot Component Scan", Category: api.Ref{Name: "Sustain"}},
		{Name: "Spring Boot Configuration", Category: api.Ref{Name: "Sustain"}},
		{Name: "Spring Boot Auto-configuration", Category: api.Ref{Name: "Configuration Management"}},
		{Name: "Spring Boot Component Scan", Category: api.Ref{Name: "Embedded"}},
	},
}
