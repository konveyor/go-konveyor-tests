package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var AcmeairWebapp = TC{
	Name:        "acmeair-webapp",
	Application: data.UploadBinary,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Binary:   true,
	Artifact: "/binary/acmeair-webapp-1.0-SNAPSHOT.war",
	Analysis: api.Analysis{
		Effort: 54,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "asm.asm",
				Version:  "3.3.1",
				Provider: "java",
				SHA: "1d5f20b4ea675e6fab6ab79f1cd60ec268ddc015",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-beans",
				Version:  "3.1.2.RELEASE",
				SHA:      "38ff7985f572a87d72023eedce3d720ce1ba4083",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-asm",
				Version:  "3.1.2.RELEASE",
				SHA:      "d45753d8048317099395ce7f2bbcdb1776a2e26c",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.aspectj.aspectjweaver",
				Version:  "1.6.8",
				SHA:      "d3fdabe0a1f7f395d64bf9d0c1a3e0f1b98430df",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-web",
				Version:  "3.1.2.RELEASE",
				SHA:      "34a73557b66ea3c68892c042f2c9ba7e8bff8eac",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "net.wasdev.wlp.sample.acmeair-common",
				Version:  "1.0-SNAPSHOT",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-core",
				Version:  "3.1.2.RELEASE",
				SHA:      "dd4295f0567deb2cc629dd647d2f055268c2fd3e",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-tx",
				Version:  "3.1.2.RELEASE",
				SHA:      "f9efe9e07e73ab593de0353e01b8bcbaafbfd3c1",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "net.wasdev.wlp.sample.acmeair-services",
				Version:  "1.0-SNAPSHOT",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
			   Provider: "java",
               Name: "net.wasdev.wlp.sample.acmeair-services-jpa",
               Version: "1.0-SNAPSHOT",
			   Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "cglib.cglib",
				Version:  "2.2.2",
				SHA:      "a47a971686474124562bdd4a7ccbd8ac8c3e8b11",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-context",
				Version:  "3.1.2.RELEASE",
				SHA:      "06787712741427c1014aaca503d24e3b3f05f396",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-expression",
				Version:  "3.1.2.RELEASE",
				SHA:      "c6d4aca58f993bda3cce2bc0536bd168cc9d7ed5",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.springframework.spring-aop",
				Version:  "3.1.2.RELEASE",
				SHA:      "a68c5ee96d735ac68c00e226e768ef70ae1d5c57",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.aspectj.aspectjrt",
				Version:  "1.6.8",
				SHA:      "8c234df6a43dd5ba355803122e33caa7ccd11011",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "aopalliance.aopalliance",
				Version:  "1.0",
				SHA:      "0235ba8b489512805ac13a8f9ea77a1ca5ebe3e8",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "commons-logging.commons-logging",
				Version:  "1.1.1",
				SHA:      "5043bfebc3db072ed80fbd362e7caf00e885d8ae",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
	},
}
