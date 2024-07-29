package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var AdministracionEfectivo = TC{
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
		Issues: []api.Issue{
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
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "EAR Deployment", Category: api.Ref{Name: "Other"}},
		{Name: "Mail", Category: api.Ref{Name: "Other"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Processing"}},
		{Name: "Mail", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "EAR Deployment", Category: api.Ref{Name: "Connect"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Execute"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Embedded"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "Mail", Category: api.Ref{Name: "Java EE"}},
		{Name: "EAR Deployment", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
	},
}
