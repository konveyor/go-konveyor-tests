package analysis

import (
	"github.com/konveyor/tackle2-hub/api"
)

var BookServerVerified = TC{
	Name: "book-server_deps",
	Task: Analyze,
	Analysis: api.Analysis{
		Effort: 15,
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/model/BookModel.java",
						Line:    14,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/model/BookModel.java",
						Line:    17,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/model/BookModel.java",
						Line:    29,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/shared/source/book-server/src/main/java/com/telran/application/model/BookModel.java",
						Line:    32,
						Message: "An application running inside a container could lose access to a file in local storage.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Spring Web artifact with Quarkus 'spring-web' extension",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "springboot-web-to-quarkus-00000",
				Incidents: []api.Incident{
					{
						File: "/shared/source/book-server/pom.xml",
						Line: 24,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Adopt Quarkus BOM",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00010",
			},
			{
				Category:    "mandatory",
				Description: "Adopt Quarkus Maven plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00020",
			},
			{
				Category:    "mandatory",
				Description: "Adopt Maven Compiler plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00030",
			},
			{
				Category:    "mandatory",
				Description: "Adopt Maven Surefire plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00040",
			},
			{
				Category:    "mandatory",
				Description: "Adopt Maven Failsafe plugin",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00050",
			},
			{
				Category:    "mandatory",
				Description: "Add Maven profile to run the Quarkus native build",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "javaee-pom-to-quarkus-00060",
			},
			{
				Category:    "mandatory",
				Description: "Remove the SpringBoot @SpringBootApplication annotation",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "springboot-annotations-to-quarkus-00000",
			},
			{
				Category:    "mandatory",
				Description: "Replace the spring-boot-maven-plugin dependency",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "springboot-plugins-to-quarkus-0000",
			},
			{
				Category:    "mandatory",
				Description: "Replace the SpringBoot artifact with Quarkus 'spring-boot-properties' extension",
				Effort:      1,
				RuleSet:     "quarkus/springboot",
				Rule:        "springboot-properties-to-quarkus-00000",
			},
		},
	},
}
