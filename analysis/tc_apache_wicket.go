package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var ApacheWicket = TC{
	Name: "Apache Wicket",
	Application: data.ApacheWicket,
	Task: Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Analysis: api.Analysis{
		Effort: 10,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Effort: 1,
				Description: "Hardcoded IP Address",
				RuleSet: "discovery-rules",
				Rule: "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup-sample-apps/test-files/src_example/org/apache/wicket/protocol/http/mock/MockHttpServletRequest.java",
						Line:    51,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
					{
						File:    "/shared/source/windup-sample-apps/test-files/src_example/org/apache/wicket/protocol/http/mock/MockHttpServletRequest.java",
						Line:    584,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
					{
						File:    "/shared/source/windup-sample-apps/test-files/src_example/org/apache/wicket/protocol/http/mock/MockHttpServletRequest.java",
						Line:    587,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
				},
			},
			{
				Category: "mandatory",
				Effort: 7,
				Description: "Local HTTP Calls",
				RuleSet: "cloud-readiness",
				Rule: "localhost-http-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup-sample-apps/test-files/src_example/org/apache/wicket/protocol/http/mock/MockHttpServletRequest.java",
						Line:    362,
						Message: "The app is trying to access local resource by HTTP, please try to migrate the resource to cloud",
					},
				},
			},
		},
	},
}
