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
		Effort: 1570,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Java native libraries (JNI, JNA)",
				Effort:      7,
				RuleSet:     "cloud-readiness",
				Rule:        "jni-native-code-00000",
			},
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
			},
			{
				Category:    "mandatory",
				Description: "File system - java.net.URL/URI",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00002",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "EAR Deployment", Category: api.Ref{Name: "Other"}},
		{Name: "JNI", Category: api.Ref{Name: "Other"}},
		{Name: "Mail", Category: api.Ref{Name: "Other"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Processing"}},
		{Name: "Mail", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "EAR Deployment", Category: api.Ref{Name: "Connect"}},
		{Name: "JNI", Category: api.Ref{Name: "Connect"}},
		{Name: "JNI", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "EAR Deployment", Category: api.Ref{Name: "Java EE"}},
		{Name: "Mail", Category: api.Ref{Name: "Java EE"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Embedded"}},
		{Name: "Spring Scheduled", Category: api.Ref{Name: "Execute"}},
	},
}
