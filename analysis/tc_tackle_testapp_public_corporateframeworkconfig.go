package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPublicCorpFrameworkConfig = TC{
	Name:        "Tackle Testapp public corporate-framework-config",
	Application: TackleTestApp,
	Task:        Analyze,
	WithDeps:    false,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=linux",
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/source=java",
		},
	},
	CustomRules: []api.RuleSet{
		{
			Name:   "This ruleset provides rules related to the corporate configuration frameworks.",
			Rules: []api.Rule{
				{
					File: &api.Ref{
						Name: "./data/corporate-framework-config.windup.xml",
					},
				},
			},
		},
	},
	Analysis: api.Analysis{
		Effort: 2,
		Issues: []api.Issue{
			{
				Category: "mandatory",
				Description: "Hardcoded IP Address\nWhen migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
				Effort: 1,
				RuleSet: "discovery-rules",
				Rule: "hardcoded-ip-address",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/src/main/resources/persistence.properties",
						Line: 2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
					{
						File: "/addon/source/tackle-testapp-public/target/classes/persistence.properties",
						Line: 2,
						Message: "When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
					},
				},
			},
			{
				Category: "potential",
				Description: "JBoss API reference",
				Effort: 0,
				RuleSet: "eap6/resin",
				Rule: "generic-catchall-00700",
				Incidents: []api.Incident{
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/controller/CustomerController.java",
						Line: 3,
						Message: "`org.jboss..logging.Logger;import io` reference found. No specific details available.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/exception/handler/ExceptionHandlingController.java",
						Line: 4,
						Message: "`org.jboss..logging.Logger;import io` reference found. No specific details available.",
					},
					{
						File: "/addon/source/tackle-testapp-public/src/main/java/io/konveyor/demo/ordermanagement/service/CustomerService.java",
						Line: 4,
						Message: "`org.jboss..logging.Logger;import io` reference found. No specific details available.",
					},
				},
			},
			{
				//TBD custom rule's issues
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Processing"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Java EE"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "Java EE Batch", Category: api.Ref{Name: "Execute"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
	},
}
