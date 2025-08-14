package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var AdministracionEfectivoBinary = TC{
	SkipTest: SkipTestConfig{
		Reason: "Skip binary test. https://issues.redhat.com/browse/MTA-5588",
		Skip:   true,
	},
	Name:        "administracion-efectivo",
	Application: data.UploadBinary,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Binary:   true,
	Artifact: "/binary/administracion_efectivo.ear",
	Analysis: api.Analysis{
		Effort: 36,
		Insights: []api.Insight{
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/export/XMLExporter.java",
						Line:     28,
						CodeSnip: "PrintWriter writer = new PrintWriter(osw);",
					},
					{
						File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/imagecropper/ImageCropperRenderer.java",
						Line:     145,
						CodeSnip: `outputImage = ImageIO.read(new File(servletContext.getRealPath("") + imagePath));`,
					},
					{
						File:     "/shared/bin/java-project/src/main/java/org/primefaces/webapp/filter/FileUploadFilter.java",
						Line:     48,
						CodeSnip: "diskFileItemFactory.setRepository(new File(this.uploadDir));",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "File system - java.net.URL/URI",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00002",
				Incidents: []api.Incident{
					{
						File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/captcha/Captcha.java",
						Line:     95,
						CodeSnip: `URL url = new URL("http://api-verify.recaptcha.net/verify");`,
					},
					{
						File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/feedreader/FeedInput.java",
						Line:     14,
						CodeSnip: "URL feedSource = new URL(url);",
					},
					{
						File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/imagecropper/ImageCropperRenderer.java",
						Line:     141,
						CodeSnip: "URL url = new URL(imagePath);",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Provider: "java",
				Name:     "jfree.jcommon",
				Version:  "1.0.15",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "6495e4f3777a24fd34c1544a8606243a600365c3",
			},
			{
				Provider: "java",
				Name:     "org.slf4j.slf4j-api",
				Version:  "1.6.1",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "6f3b8a24bf970f17289b234284c94f43eb42f0e4",
			},
			{
				Provider: "java",
				Name:     "org.bouncycastle.bctsp-jdk14",
				Version:  "1.38",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "4821122f8390d15f4b5ee652621e2a2bb1f1bf16",
			},
			{
				Provider: "java",
				Name:     "org.javassist.javassist",
				Version:  "3.15.0-GA",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "79907309ca4bb4e5e51d4086cc4179b2611358d7",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.jackson.core.jackson-databind",
				Version:  "2.1.4",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "61d1e5aaf474462b0c9bfef8aeb6a7205bd56a90",
			},
			{
				Provider: "java",
				Name:     "commons-beanutils.commons-beanutils",
				Version:  "1.8.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "0c651d5103c649c12b20d53731643e5fffceb536",
			},
			{
				Provider: "java",
				Name:     "ojdbc6-11.2.0.3.jar",
			},
			{
				Provider: "java",
				Name:     "AdministracionEfectivoGrupo.AdministracionEfectivo-ejb",
				Version:  "0.0.1-SNAPSHOT",
				Labels: []string{
					"konveyor.io/dep-source=internal",
					"konveyor.io/language=java",
				},
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-asm",
				Version:  "3.0.5.RELEASE",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "07f22a0e9f325e6565b4ea56b479ad76311d146b",
			},
			{
				Provider: "java",
				Name:     "bcprov-jdk14-1.38.jar",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.hibernate-validator",
				Version:  "4.1.0.Final",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "785cceeb0b0dbd03367f45cd60eb26cd48167640",
			},
			{
				Provider: "java",
				Name:     "j2ee.jar",
			},
			{
				Provider: "java",
				Name:     "org.codehaus.castor.castor",
				Version:  "1.2",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "351847d8a41175565b5faa54fab42fe20fe6cc91",
			},
			{
				Provider: "java",
				Name:     "org.jboss.logging.jboss-logging",
				Version:  "3.1.0.CR2",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "28725380c07f917ace4e511db21cc45e9ae5a72b",
			},
			{
				Provider: "java",
				Name:     "antlr.antlr",
				Version:  "2.7.7",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "83cd2cd674a217ade95a4bb83a8a14f351f48bd0",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-aop",
				Version:  "3.0.5.RELEASE",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "0c7a17803cc10512e26e285073639543f0c7c764",
			},
			{
				Provider: "java",
				Name:     "commons-digester.commons-digester",
				Version:  "2.1",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "73a8001e7a54a255eef0f03521ec1805dc738ca0",
			},
			{
				Provider: "java",
				Name:     "org.slf4j.slf4j-log4j12",
				Version:  "1.7.5",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "6edffc576ce104ec769d954618764f39f0f0f10d",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-beans",
				Version:  "3.0.5.RELEASE",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "4b352a9c3b427294e264ca4d460d07417ca9350e",
			},
			{
				Provider: "java",
				Name:     "c3p0.c3p0",
				Version:  "0.9.1.1",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "302704f30c6e7abb7a0457f7771739e03c973e80",
			},
			{
				Provider: "java",
				Name:     "bcmail-jdk14-138.jar",
			},
			{
				Provider: "java",
				Name:     "bcmail-jdk14-1.38.jar",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-context",
				Version:  "3.0.5.RELEASE",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "6b05e397566cc7750d2d25f81a7441fe1aeecb75",
			},
			{
				Provider: "java",
				Name:     "itext-2.1.7.js2.jar",
			},
			{
				Provider: "java",
				Name:     "commons-logging.commons-logging",
				Version:  "1.1.1",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "5043bfebc3db072ed80fbd362e7caf00e885d8ae",
			},
			{
				Provider: "java",
				Name:     "commons-collections.commons-collections",
				Version:  "2.1",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "8e10f94f30ea064eee3cb94f864dc9c31e30e8af",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.jackson.core.jackson-annotations",
				Version:  "2.1.4",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "76b956e45e19e4e081a9c52a68a3c078e3b2cf17",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.javax.persistence.hibernate-jpa-2.0-api",
				Version:  "1.0.1.Final",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "3306a165afa81938fc3d8a0948e891de9f6b192b",
			},
			{
				Provider: "java",
				Name:     "org.jasypt.jasypt",
				Version:  "1.9.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "0857a1a55a81641c31b2a9b4b292120c1d4432bd",
			},
			{
				Provider: "java",
				Name:     "org.primefaces.primefaces",
				Version:  "3.5",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
			},
			{
				Provider: "java",
				Name:     "org.eclipse.persistence.javax.persistence",
				Version:  "2.0.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "bff9b1d9de629095001f1a4e77f450b2d6487b07",
			},
			{
				Provider: "java",
				Name:     "com.ibm.ws.jsf.jar",
			},
			{
				Provider: "java",
				Name:     "org.jboss.spec.javax.transaction.jboss-transaction-api_1.1_spec",
				Version:  "1.0.0.Final",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "2ab6236535e085d86f37fd97ddfdd35c88c1a419",
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-expression",
				Version:  "3.0.5.RELEASE",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "5b8e53877cb58c94f15a0d8172da3569f4b4f3fb",
			},
			{
				Provider: "java",
				Name:     "eclipse.jdtcore",
				Version:  "3.1.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "c5e3e72ae7220118c3da808628ec7016d4d8aef2",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.hibernate-core",
				Version:  "4.1.4.Final",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "811b5e2b3e13adae69078f8435990812276ea9df",
			},
			{
				Provider: "java",
				Name:     "dom4j-1.6.1.jar",
			},
			{
				Provider: "java",
				Name:     "aopalliance.aopalliance",
				Version:  "1.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "0235ba8b489512805ac13a8f9ea77a1ca5ebe3e8",
			},
			{
				Provider: "java",
				Name:     "net.sf.jasperreports.jasperreports",
				Version:  "5.5.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "4eb362300120b5cbf0cdd9f0c59c7db2bcfdef65",
			},
			{
				Provider: "java",
				Name:     "bcprov-jdk14-138.jar",
			},
			{
				Provider: "java",
				Name:     "javax.validation.validation-api",
				Version:  "1.0.0.GA",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "b6bd7f9d78f6fdaa3c37dae18a4bd298915f328e",
			},
			{
				Provider: "java",
				Name:     "AdministracionEfectivoGrupo.AdministracionEfectivo-seguridad",
				Version:  "0.0.1-SNAPSHOT",
				Labels: []string{
					"konveyor.io/dep-source=internal",
					"konveyor.io/language=java",
				},
			},
			{
				Provider: "java",
				Name:     "org.springframework.spring-core",
				Version:  "3.0.5.RELEASE",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "1633e94943d57746ef76910489f1cd71fe667e04",
			},
			{
				Provider: "java",
				Name:     "com.fasterxml.jackson.core.jackson-core",
				Version:  "2.1.4",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "08896e10f8b07918b60c835b45a05507fa213e84",
			},
			{
				Provider: "java",
				Name:     "org.quartz-scheduler.quartz",
				Version:  "2.2.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "2eb16fce055d5f3c9d65420f6fc4efd3a079a3d8",
			},
			{
				Provider: "java",
				Name:     "jfree.jfreechart",
				Version:  "1.0.12",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "47e8a2e9e052d292e416a6fd5292a77b54c48fac",
			},
			{
				Provider: "java",
				Name:     "log4j.log4j",
				Version:  "1.2.12",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "057b8740427ee6d7b0b60792751356cad17dc0d9",
			},
			{
				Provider: "java",
				Name:     "AdministracionEfectivoGrupo.AdministracionEfectivo-jpa",
				Version:  "0.0.1-SNAPSHOT",
				Labels: []string{
					"konveyor.io/dep-source=internal",
					"konveyor.io/language=java",
				},
			},
			{
				Provider: "java",
				Name:     "xml-apis.xml-apis",
				Version:  "1.3.02",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "dc7315e359b5e43f20131414d60b5c307aace975",
			},
			{
				Provider: "java",
				Name:     "org.quartz-scheduler.quartz-jobs",
				Version:  "2.2.0",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "865fa48d18a3fee0164a14423f0fbef3dc5fd04b",
			},
			{
				Provider: "java",
				Name:     "org.hibernate.common.hibernate-commons-annotations",
				Version:  "4.0.1.Final",
				Labels: []string{
					"konveyor.io/dep-source=open-source",
					"konveyor.io/language=java",
				},
				SHA: "78bcf608d997d0529be2f4f781fdc89e801c9e88",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "GNU LGPL", Category: api.Ref{Name: "License"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Mail", Category: api.Ref{Name: "Other"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Processing"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Processing"}},
		{Name: "Mail", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Mail", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Embedded"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Execute"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Execute"}},
		{Name: "JSF XML", Category: api.Ref{Name: "View"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Web"}},
		{Name: "JSF Page", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF Page", Category: api.Ref{Name: "View"}},
		{Name: "JSF Page", Category: api.Ref{Name: "Web"}},
		{Name: "JSF", Category: api.Ref{Name: "MVC"}},
		{Name: "JSF", Category: api.Ref{Name: "View"}},
		{Name: "JSF", Category: api.Ref{Name: "Embedded"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Java EE"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Other"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Connect"}},
	},
}
