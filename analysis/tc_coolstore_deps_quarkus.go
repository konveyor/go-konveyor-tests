package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/data/identity"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var CoolstoreDepsQuarkus = TC{
	Name:        "Coolstore Source and Deps analysis using Quarkus branch",
	Application: data.CoolstoreQuarkus,
	WithDeps: true,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=jakarta-ee",
		},
	},
	Identities: []api.Identity{
		identity.TackleTestappPublicMaven,
	},
	Analysis: api.Analysis{
		Effort: 7,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "Local HTTP Calls",
				Effort: 7,
				RuleSet: "cloud-readiness",
				Rule: "localhost-http-00001",
				Incidents: []api.Incident{
					{
						File: "/shared/source/coolstore/src/main/resources/application.properties",
						Line: 10,
						Message: "The app is trying to access local resource by HTTP, please try to migrate the resource to cloud",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name: "io.quarkus.quarkus-jdbc-postgresql",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-kubernetes",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-kubernetes-client-internal",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-messaging",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-minikube",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-flyway",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-resteasy-client",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-resteasy-client-jackson",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-resteasy-jackson",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-undertow",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "org.flywaydb.flyway-database-postgresql",
				Version: "10.12.0",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-hibernate-orm",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-container-image-docker",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-arc",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-resteasy",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-expression",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-mutiny-vertx-core",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "com.github.java-json-tools.json-patch",
				Version: "1.13",
				Provider: "java",
			},
			{
				Name: "com.github.java-json-tools.msg-simple",
				Version: "1.2",
				Provider: "java",
			},
			{
				Name: "com.google.code.gson.gson",
				Version: "2.11.0",
				Provider: "java",
			},
			{
				Name: "com.google.errorprone.error_prone_annotations",
				Version: "2.28.0",
				Provider: "java",
			},
			{
				Name: "com.ibm.async.asyncutil",
				Version: "0.1.0",
				Provider: "java",
			},
			{
				Name: "com.sun.istack.istack-commons-runtime",
				Version: "4.1.2",
				Provider: "java",
			},
			{
				Name: "commons-codec.commons-codec",
				Version: "1.17.0",
				Provider: "java",
			},
			{
				Name: "io.agroal.agroal-api",
				Version: "2.4",
				Provider: "java",
			},
			{
				Name: "io.agroal.agroal-narayana",
				Version: "2.4",
				Provider: "java",
			},
			{
				Name: "io.agroal.agroal-pool",
				Version: "2.4",
				Provider: "java",
			},
			{
				Name: "io.github.crac.org-crac",
				Version: "0.1.3",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-buffer",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-codec",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-codec-dns",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-codec-haproxy",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-codec-http",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-codec-http2",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-codec-socks",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-common",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-handler",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-handler-proxy",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-resolver",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-resolver-dns",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-transport",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.netty.netty-transport-native-unix-common",
				Version: "4.1.108.Final",
				Provider: "java",
			},
			{
				Name: "io.opentelemetry.opentelemetry-api",
				Version: "1.32.0",
				Provider: "java",
			},
			{
				Name: "io.opentelemetry.opentelemetry-context",
				Version: "1.32.0",
				Provider: "java",
			},
			{
				Name: "io.quarkus.arc.arc",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.http.quarkus-http-core",
				Version: "5.2.2.Final",
				Provider: "java",
			},
			{
				Name: "io.quarkus.http.quarkus-http-http-core",
				Version: "5.2.2.Final",
				Provider: "java",
			},
			{
				Name: "io.quarkus.http.quarkus-http-servlet",
				Version: "5.2.2.Final",
				Provider: "java",
			},
			{
				Name: "io.quarkus.http.quarkus-http-vertx-backend",
				Version: "5.2.2.Final",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-agroal",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-apache-httpclient",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.github.java-json-tools.btf",
				Version: "1.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-caffeine",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-classloader-commons",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-container-image",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.github.ben-manes.caffeine.caffeine",
				Version: "3.1.5",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-container-image-docker-common",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-core",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-credentials",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-datasource",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-datasource-common",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-development-mode-spi",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.module.jackson-module-parameter-names",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-fs-util",
				Version: "0.0.10",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.module.jackson-module-jakarta-xmlbind-annotations",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-ide-launcher",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-jackson",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.jakarta.rs.jackson-jakarta-rs-json-provider",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.jakarta.rs.jackson-jakarta-rs-base",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.datatype.jackson-datatype-jsr310",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.datatype.jackson-datatype-jdk8",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-messaging-kotlin",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.dataformat.jackson-dataformat-toml",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-mutiny",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-mutiny-reactive-streams-operators",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-narayana-jta",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-netty",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-rest-client-config",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-databind",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-core",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.jackson.core.jackson-annotations",
				Version: "2.17.2",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-resteasy-common",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.fasterxml.classmate",
				Version: "1.7.0",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-resteasy-server-common",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-security-runtime-spi",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-smallrye-context-propagation",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-tls-registry",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-transaction-annotations",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "com.aayushatharva.brotli4j.service",
				Version: "1.16.0",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-vertx",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-vertx-http",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-vertx-latebound-mdc-provider",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.quarkus-virtual-threads",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.quarkus.security.quarkus-security",
				Version: "2.1.0",
				Provider: "java",
			},
			{
				Name: "io.quarkus.vertx.utils.quarkus-vertx-utils",
				Version: "3.12.3",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-annotation",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-classloader",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-constraint",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-cpu",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "com.aayushatharva.brotli4j.brotli4j",
				Version: "1.16.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-function",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-io",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-net",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-os",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-ref",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.common.smallrye-common-vertx-context",
				Version: "2.3.0",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.resteasy-client-api",
				Version: "6.2.9.Final",
				Provider: "java",
			},
			{
				Name: "com.github.java-json-tools.jackson-coreutils",
				Version: "2.0",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.microprofile.microprofile-rest-client-base",
				Version: "2.1.5.Final",
				Provider: "java",
			},
			{
				Name: "io.smallrye.jandex",
				Version: "3.2.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.mutiny",
				Version: "2.6.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.mutiny-reactive-streams-operators",
				Version: "2.6.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.mutiny-smallrye-context-propagation",
				Version: "2.6.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.mutiny-zero",
				Version: "1.1.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.mutiny-zero-flow-adapters",
				Version: "1.1.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-mutiny-vertx-auth-common",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-mutiny-vertx-bridge-common",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.config.smallrye-config-common",
				Version: "3.8.3",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-mutiny-vertx-runtime",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-mutiny-vertx-uri-template",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-mutiny-vertx-web",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-mutiny-vertx-web-common",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-reactive-converter-api",
				Version: "3.0.1",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-reactive-converter-mutiny",
				Version: "3.0.1",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-reactive-messaging-api",
				Version: "4.21.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-reactive-messaging-health",
				Version: "4.21.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.smallrye-reactive-messaging-provider",
				Version: "4.21.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.reactive.vertx-mutiny-generator",
				Version: "3.12.0",
				Provider: "java",
			},
			{
				Name: "io.smallrye.smallrye-context-propagation",
				Version: "2.1.2",
				Provider: "java",
			},
			{
				Name: "io.smallrye.smallrye-context-propagation-api",
				Version: "2.1.2",
				Provider: "java",
			},
			{
				Name: "io.smallrye.smallrye-context-propagation-jta",
				Version: "2.1.2",
				Provider: "java",
			},
			{
				Name: "io.smallrye.smallrye-context-propagation-storage",
				Version: "2.1.2",
				Provider: "java",
			},
			{
				Name: "io.smallrye.smallrye-fault-tolerance-vertx",
				Version: "6.3.0",
				Provider: "java",
			},
			{
				Name: "io.vertx.vertx-auth-common",
				Version: "4.5.7",
				Provider: "java",
			},
			{
				Name: "io.vertx.vertx-bridge-common",
				Version: "4.5.7",
				Provider: "java",
			},
			{
				Name: "io.vertx.vertx-codegen",
				Version: "4.5.7",
				Provider: "java",
			},
			{
				Name: "io.vertx.vertx-core",
				Version: "4.5.7",
				Provider: "java",
			},
			{
				Name: "io.vertx.vertx-uri-template",
				Version: "4.5.7",
				Provider: "java",
			},
			{
				Name: "io.vertx.vertx-web",
				Version: "4.5.7",
				Provider: "java",
			},
			{
				Name: "io.vertx.vertx-web-common",
				Version: "4.5.7",
				Provider: "java",
			},
			{
				Name: "jakarta.activation.jakarta.activation-api",
				Version: "2.1.3",
				Provider: "java",
			},
			{
				Name: "jakarta.annotation.jakarta.annotation-api",
				Version: "3.0.0",
				Provider: "java",
			},
			{
				Name: "jakarta.el.jakarta.el-api",
				Version: "5.0.1",
				Provider: "java",
			},
			{
				Name: "jakarta.enterprise.jakarta.enterprise.cdi-api",
				Version: "4.1.0",
				Provider: "java",
			},
			{
				Name: "jakarta.enterprise.jakarta.enterprise.lang-model",
				Version: "4.1.0",
				Provider: "java",
			},
			{
				Name: "jakarta.inject.jakarta.inject-api",
				Version: "2.0.1",
				Provider: "java",
			},
			{
				Name: "jakarta.interceptor.jakarta.interceptor-api",
				Version: "2.2.0",
				Provider: "java",
			},
			{
				Name: "jakarta.json.jakarta.json-api",
				Version: "2.1.3",
				Provider: "java",
			},
			{
				Name: "jakarta.persistence.jakarta.persistence-api",
				Version: "3.1.0",
				Provider: "java",
			},
			{
				Name: "jakarta.resource.jakarta.resource-api",
				Version: "2.1.0",
				Provider: "java",
			},
			{
				Name: "jakarta.servlet.jakarta.servlet-api",
				Version: "6.0.0",
				Provider: "java",
			},
			{
				Name: "jakarta.transaction.jakarta.transaction-api",
				Version: "2.0.1",
				Provider: "java",
			},
			{
				Name: "jakarta.validation.jakarta.validation-api",
				Version: "3.0.2",
				Provider: "java",
			},
			{
				Name: "jakarta.ws.rs.jakarta.ws.rs-api",
				Version: "3.1.0",
				Provider: "java",
			},
			{
				Name: "jakarta.xml.bind.jakarta.xml.bind-api",
				Version: "4.0.2",
				Provider: "java",
			},
			{
				Name: "net.bytebuddy.byte-buddy",
				Version: "1.14.15",
				Provider: "java",
			},
			{
				Name: "org.antlr.antlr4-runtime",
				Version: "4.13.0",
				Provider: "java",
			},
			{
				Name: "org.apache.httpcomponents.httpasyncclient",
				Version: "4.1.5",
				Provider: "java",
			},
			{
				Name: "org.apache.httpcomponents.httpclient",
				Version: "4.5.14",
				Provider: "java",
			},
			{
				Name: "org.apache.httpcomponents.httpcore",
				Version: "4.4.16",
				Provider: "java",
			},
			{
				Name: "org.apache.httpcomponents.httpcore-nio",
				Version: "4.4.16",
				Provider: "java",
			},
			{
				Name: "org.eclipse.angus.angus-activation",
				Version: "2.0.2",
				Provider: "java",
			},
			{
				Name: "org.eclipse.microprofile.config.microprofile-config-api",
				Version: "3.1",
				Provider: "java",
			},
			{
				Name: "org.eclipse.microprofile.context-propagation.microprofile-context-propagation-api",
				Version: "1.3",
				Provider: "java",
			},
			{
				Name: "org.eclipse.microprofile.health.microprofile-health-api",
				Version: "4.0.1",
				Provider: "java",
			},
			{
				Name: "org.eclipse.microprofile.reactive-streams-operators.microprofile-reactive-streams-operators-api",
				Version: "3.0",
				Provider: "java",
			},
			{
				Name: "org.eclipse.microprofile.reactive-streams-operators.microprofile-reactive-streams-operators-core",
				Version: "3.0",
				Provider: "java",
			},
			{
				Name: "org.eclipse.microprofile.rest.client.microprofile-rest-client-api",
				Version: "3.0.1",
				Provider: "java",
			},
			{
				Name: "org.eclipse.parsson.parsson",
				Version: "1.1.6",
				Provider: "java",
			},
			{
				Name: "org.flywaydb.flyway-core",
				Version: "10.15.2",
				Provider: "java",
			},
			{
				Name: "com.aayushatharva.brotli4j.native-linux-x86_64",
				Version: "1.16.0",
				Provider: "java",
			},
			{
				Name: "org.glassfish.jaxb.jaxb-core",
				Version: "4.0.5",
				Provider: "java",
			},
			{
				Name: "org.glassfish.jaxb.jaxb-runtime",
				Version: "4.0.5",
				Provider: "java",
			},
			{
				Name: "org.glassfish.jaxb.txw2",
				Version: "4.0.5",
				Provider: "java",
			},
			{
				Name: "org.hibernate.common.hibernate-commons-annotations",
				Version: "6.0.6.Final",
				Provider: "java",
			},
			{
				Name: "org.hibernate.orm.hibernate-core",
				Version: "6.5.2.Final",
				Provider: "java",
			},
			{
				Name: "org.hibernate.orm.hibernate-graalvm",
				Version: "6.5.2.Final",
				Provider: "java",
			},
			{
				Name: "org.hibernate.quarkus-local-cache",
				Version: "0.3.0",
				Provider: "java",
			},
			{
				Name: "org.jboss.invocation.jboss-invocation",
				Version: "2.0.0.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.jboss-transaction-spi",
				Version: "8.0.0.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.logging.commons-logging-jboss-logging",
				Version: "1.0.0.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.logging.jboss-logging",
				Version: "3.6.0.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.logging.jboss-logging-annotations",
				Version: "2.2.1.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.logmanager.jboss-logmanager",
				Version: "3.0.6.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.narayana.jta.narayana-jta",
				Version: "7.0.2.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.narayana.jts.narayana-jts-integration",
				Version: "7.0.2.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.microprofile.microprofile-config",
				Version: "2.1.5.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.microprofile.microprofile-rest-client",
				Version: "2.1.5.Final",
				Provider: "java",
			},
			{
				Name: "io.smallrye.config.smallrye-config-core",
				Version: "3.8.3",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.resteasy-cdi",
				Version: "6.2.9.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.resteasy-client",
				Version: "6.2.9.Final",
				Provider: "java",
			},
			{
				Name: "io.smallrye.config.smallrye-config",
				Version: "3.8.3",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.resteasy-core",
				Version: "6.2.9.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.resteasy-core-spi",
				Version: "6.2.9.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.resteasy.resteasy-jackson2-provider",
				Version: "6.2.9.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.slf4j.slf4j-jboss-logmanager",
				Version: "2.0.0.Final",
				Provider: "java",
			},
			{
				Name: "org.jboss.threads.jboss-threads",
				Version: "3.6.1.Final",
				Provider: "java",
			},
			{
				Name: "org.jctools.jctools-core",
				Version: "4.0.3",
				Provider: "java",
			},
			{
				Name: "org.postgresql.postgresql",
				Version: "42.7.3",
				Provider: "java",
			},
			{
				Name: "org.reactivestreams.reactive-streams",
				Version: "1.0.4",
				Provider: "java",
			},
			{
				Name: "org.slf4j.slf4j-api",
				Version: "2.0.6",
				Provider: "java",
			},
			{
				Name: "org.wildfly.common.wildfly-common",
				Version: "1.7.0.Final",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "CDI", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "RMI", Category: api.Ref{Name: "Other"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "Apache License 2.0", Category: api.Ref{Name: "License"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "CDI", Category: api.Ref{Name: "Java EE"}},
		{Name: "RMI", Category: api.Ref{Name: "Java EE"}},
		{Name: "GNU GPL", Category: api.Ref{Name: "License"}},
		{Name: "Application Properties File", Category: api.Ref{Name: "Embedded"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Application Properties File", Category: api.Ref{Name: "Configuration Management"}},
		{Name: "CDI", Category: api.Ref{Name: "Execute"}},
		{Name: "Application Properties File", Category: api.Ref{Name: "Sustain"}},
		{Name: "RMI", Category: api.Ref{Name: "Connect"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
	},
}
