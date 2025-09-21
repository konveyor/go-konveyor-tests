package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var AdministracionEfectivoBinary = TC{
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
		Effort: 30,
		// Insights: []api.Insight{
		// 	{
		// 		Category:    "mandatory",
		// 		Description: "File system - Java IO",
		// 		Effort:      1,
		// 		RuleSet:     "cloud-readiness",
		// 		Rule:        "local-storage-00001",
		// 		Incidents: []api.Incident{
		// 			{
		// 				File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/export/XMLExporter.java",
		// 				Line:     28,
		// 				CodeSnip: "PrintWriter writer = new PrintWriter(osw);",
		// 			},
		// 			{
		// 				File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/imagecropper/ImageCropperRenderer.java",
		// 				Line:     145,
		// 				CodeSnip: `outputImage = ImageIO.read(new File(servletContext.getRealPath("") + imagePath));`,
		// 			},
		// 			{
		// 				File:     "/shared/bin/java-project/src/main/java/org/primefaces/webapp/filter/FileUploadFilter.java",
		// 				Line:     48,
		// 				CodeSnip: "diskFileItemFactory.setRepository(new File(this.uploadDir));",
		// 			},
		// 		},
		// 	},
		// 	{
		// 		Category:    "mandatory",
		// 		Description: "File system - java.net.URL/URI",
		// 		Effort:      1,
		// 		RuleSet:     "cloud-readiness",
		// 		Rule:        "local-storage-00002",
		// 		Incidents: []api.Incident{
		// 			{
		// 				File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/captcha/Captcha.java",
		// 				Line:     95,
		// 				CodeSnip: `URL url = new URL("http://api-verify.recaptcha.net/verify");`,
		// 			},
		// 			{
		// 				File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/feedreader/FeedInput.java",
		// 				Line:     14,
		// 				CodeSnip: "URL feedSource = new URL(url);",
		// 			},
		// 			{
		// 				File:     "/shared/bin/java-project/src/main/java/org/primefaces/component/imagecropper/ImageCropperRenderer.java",
		// 				Line:     141,
		// 				CodeSnip: "URL url = new URL(imagePath);",
		// 			},
		// 		},
		// 	},
		// },
		Dependencies: []api.TechDependency{
			{
				Name:     "org.eclipse.jdt.jdtcore-3.1.0",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.slf4j.slf4j-log4j12",
				Version:  "1.7.5",
				Provider: "java",
			},
			{
				Name:     "org.bouncycastle.bcmail-jdk14-138",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     ".javax.persistence-2.0.0.jar",
				Version:  "",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-context-3.0.5.RELEASE",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.apache.commons.commons-collections-2.1",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.primefaces.primefaces",
				Version:  "3.5",
				Provider: "java",
			},
			{
				Name:     "AdministracionEfectivoGrupo.AdministracionEfectivo-ejb",
				Version:  "0.0.1-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-core-3.0.5.RELEASE",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.bouncycastle.bcprov-jdk14-138",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.jasypt.jasypt",
				Version:  "1.9.0",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-annotations",
				Version:  "2.1.4",
				Provider: "java",
			},
			{
				Name:     "com.ibm.com.ibm.ws.jsf",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.jboss.spec.javax.transaction.jboss-transaction-api_1.1_spec",
				Version:  "1.0.0.Final",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-asm-3.0.5.RELEASE",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-beans-3.0.5.RELEASE",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "commons-logging.commons-logging",
				Version:  "1.1.1",
				Provider: "java",
			},
			{
				Name:     "org.quartz-scheduler.quartz",
				Version:  "2.2.0",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-expression-3.0.5.RELEASE",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     ".jcommon-1.0.15.jar",
				Version:  "",
				Provider: "java",
			},
			{
				Name:     "org.dom4j.dom4j-1.6.1",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "com.lowagie.itext-2.1.7.js2",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "AdministracionEfectivoGrupo.AdministracionEfectivo-jpa",
				Version:  "0.0.1-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     ".j2ee.jar",
				Version:  "",
				Provider: "java",
			},
			{
				Name:     "org.bouncycastle.bctsp-jdk14-1.38",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "net.sf.jasperreports.jasperreports",
				Version:  "5.5.0",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-core",
				Version:  "2.1.4",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.javax.persistence.hibernate-jpa-2.0-api",
				Version:  "1.0.1.Final",
				Provider: "java",
			},
			{
				Name:     "org.javassist.javassist",
				Version:  "3.15.0-GA",
				Provider: "java",
			},
			{
				Name:     ".antlr-2.7.7.jar",
				Version:  "",
				Provider: "java",
			},
			{
				Name:     "org.jboss.logging.jboss-logging",
				Version:  "3.1.0.CR2",
				Provider: "java",
			},
			{
				Name:     "commons-beanutils.commons-beanutils",
				Version:  "1.8.0",
				Provider: "java",
			},
			{
				Name:     "com.fasterxml.jackson.core.jackson-databind",
				Version:  "2.1.4",
				Provider: "java",
			},
			{
				Name:     ".xml-apis-1.3.02.jar",
				Version:  "",
				Provider: "java",
			},
			{
				Name:     "org.jfree.jfreechart-1.0.12",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "commons-digester.commons-digester",
				Version:  "2.1",
				Provider: "java",
			},
			{
				Name:     "org.codehaus.castor.castor",
				Version:  "1.2",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.common.hibernate-commons-annotations",
				Version:  "4.0.1.Final",
				Provider: "java",
			},
			{
				Name:     "org.slf4j.slf4j-api",
				Version:  "1.6.1",
				Provider: "java",
			},
			{
				Name:     "org.quartz-scheduler.quartz-jobs",
				Version:  "2.2.0",
				Provider: "java",
			},
			{
				Name:     "javax.validation.validation-api",
				Version:  "1.0.0.GA",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.hibernate-validator",
				Version:  "4.1.0.Final",
				Provider: "java",
			},
			{
				Name:     "org.springframework.spring-aop-3.0.5.RELEASE",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     ".aopalliance-1.0.jar",
				Version:  "",
				Provider: "java",
			},
			{
				Name:     "com.mchange.c3p0-0.9.1.1",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.bouncycastle.bcprov-jdk14-1.38",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.bouncycastle.bcmail-jdk14-1.38",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     "org.apache.log4j-1.2.12",
				Version:  "Unknown",
				Provider: "java",
			},
			{
				Name:     ".ojdbc6-11.2.0.3.jar",
				Version:  "",
				Provider: "java",
			},
			{
				Name:     "AdministracionEfectivoGrupo.AdministracionEfectivo-seguridad",
				Version:  "0.0.1-SNAPSHOT",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.hibernate-core-4.1.4.Final",
				Version:  "Unknown",
				Provider: "java",
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
