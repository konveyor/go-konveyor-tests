package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var UploadBinary = TC{
	Name:        "upload-binary",
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
		Effort: 801,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Method encodeRedirectUrl in javax.servlet.http.HttpServletResponse has been removed",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-servlet-00101",
			},
			{
				Category:    "mandatory",
				Description: "Method encodeRedirectUrl in javax.servlet.http.HttpServletResponseWrapper has been removed",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-servlet-00111",
			},
			{
				Category:    "mandatory",
				Description: "Method encodetUrl in javax.servlet.http.HttpServletResponse has been removed",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "javax-to-jakarta-servlet-00100",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
	},
}
