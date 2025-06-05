package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var SeamBooking = TC{
	Name:        "Seam booking",
	Application: data.SeamBooking,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=eap",
		},
	},
	Analysis: api.Analysis{
		Effort: 89,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Hibernate 5.3 - default_schema or default_catalog must be defined or set jdbc_metadata_extraction_strategy",
				Effort:      1,
				RuleSet:     "eap7/weblogic/tests/data",
				Rule:        "hibernate51-53-01000",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/resources/META-INF/persistence.xml",
						Line:    7,
						Message: "Define `hibernate.default_schema` or `hibernate.default_catalog` (whichever is used by the selected dialect), or, alternatively, set `hibernate.hbm2ddl.jdbc_metadata_extraction_strategy=individually`.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/resources/META-INF/persistence.xml",
						Line:    9,
						Message: "Define `hibernate.default_schema` or `hibernate.default_catalog` (whichever is used by the selected dialect), or, alternatively, set `hibernate.hbm2ddl.jdbc_metadata_extraction_strategy=individually`.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:button",
				Effort:      1,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-00001-01",
				Incidents: []api.Incident{
					{
						File: "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line: 94,
						Message: "Seam UI's `<s:button>` JSF control should be replaced by `<h:button>`.\n" +
							" There are differences in attributes, for example the _action_ attribute maps to _outcome_ and there is no _propagation_ attribute.",
					},
					{
						File: "/shared/source/windup/test-files/seam-booking-5.2/view/confirm.xhtml",
						Line: 42,
						Message: "Seam UI's `<s:button>` JSF control should be replaced by `<h:button>`.\n" +
							" There are differences in attributes, for example the _action_ attribute maps to _outcome_ and there is no _propagation_ attribute.",
					},
					{
						File: "/shared/source/windup/test-files/seam-booking-5.2/view/password.xhtml",
						Line: 36,
						Message: "Seam UI's `<s:button>` JSF control should be replaced by `<h:button>`.\n" +
							" There are differences in attributes, for example the _action_ attribute maps to _outcome_ and there is no _propagation_ attribute.",
					},
					{
						File: "/shared/source/windup/test-files/seam-booking-5.2/view/register.xhtml",
						Line: 81,
						Message: "Seam UI's `<s:button>` JSF control should be replaced by `<h:button>`.\n" +
							" There are differences in attributes, for example the _action_ attribute maps to _outcome_ and there is no _propagation_ attribute",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:convertDateTime",
				Effort:      1,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-01005",
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:decorate",
				Effort:      3,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-01023",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line:    29,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line:    34,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line:    39,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line:    48,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line:    56,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line:    63,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/book.xhtml",
						Line:    70,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/confirm.xhtml",
						Line:    21,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/confirm.xhtml",
						Line:    28,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/confirm.xhtml",
						Line:    33,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/confirm.xhtml",
						Line:    38,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/hotelview.xhtml",
						Line:    9,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/hotelview.xhtml",
						Line:    14,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/hotelview.xhtml",
						Line:    19,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/hotelview.xhtml",
						Line:    24,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/hotelview.xhtml",
						Line:    29,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/hotelview.xhtml",
						Line:    34,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/hotelview.xhtml",
						Line:    39,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/password.xhtml",
						Line:    27,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/password.xhtml",
						Line:    32,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/register.xhtml",
						Line:    58,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/register.xhtml",
						Line:    65,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/register.xhtml",
						Line:    72,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/register.xhtml",
						Line:    77,
						Message: "There is no direct mapping for <s:decorate> in JSF UI controls, but you can achieve the same functionality by using the UIInputContainer and a composite container, both of which are demonstrated in the Open18 migration example input.xhtml file.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:label",
				Effort:      1,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-01026",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/edit.xhtml",
						Line:    11,
						Message: "There is no direct mapping for <s:label> in JSF UI controls, but <h:outputLabel> is similar.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:link",
				Effort:      1,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-01000",
				Incidents: []api.Incident{
					{
						File: "/shared/source/windup/test-files/seam-booking-5.2/view/home.xhtml",
						Line: 10,
						Message: "Seam UI's `<s:link>` should be replaced by `<h:link>`.\n" +
							" There are differences in attributes, for example the _action_ attribute maps to _outcome_ and there is no _propagation_ attribute.",
					},
					{
						File: "/shared/source/windup/test-files/seam-booking-5.2/view/main.xhtml",
						Line: 68,
						Message: "Seam UI's `<s:link>` should be replaced by `<h:link>`.\n" +
							" There are differences in attributes, for example the _action_ attribute maps to _outcome_ and there is no _propagation_ attribute.",
					},
					{
						File: "/shared/source/windup/test-files/seam-booking-5.2/view/template.xhtml",
						Line: 13,
						Message: "Seam UI's `<s:link>` should be replaced by `<h:link>`.\n" +
							" There are differences in attributes, for example the _action_ attribute maps to _outcome_ and there is no _propagation_ attribute.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:message",
				Effort:      1,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-01027",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/edit.xhtml",
						Line:    18,
						Message: `Use <h:message for="name" errorClass="invalid"> or Richfaces <rich:message>.`,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:span",
				Effort:      1,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-01025",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/edit.xhtml",
						Line:    11,
						Message: "There is no direct mapping for <s:span> in JSF UI controls, but you can achieve a similar effect by using <h:panelGroup> or a <ui:fragment> with a span element.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "JSF Seam 2 UI control s:validateAll",
				Effort:      1,
				RuleSet:     "eap6/java-ee/seam",
				Rule:        "seam-ui-jsf-01022",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/windup/test-files/seam-booking-5.2/view/edit.xhtml",
						Line:    14,
						Message: "There is no direct mapping for <s:validateAll> in JSF UI controls, but you can achieve a similar effect by using <f:validateBean> or Richfaces <rich:validator>.	",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "com.fasterxml.classmate",
				Version:  "1.3.4",
				Provider: "java",
			},
			{
				Name:     "javax.persistence.persistence-api",
				Version:  "1.0",
				Provider: "java",
			},
			{
				Name:     "javax.validation.validation-api",
				Version:  "2.0.1.Final",
				Provider: "java",
			},
			{
				Name:     "org.hibernate.validator.hibernate-validator",
				Version:  "6.0.23.Final",
				Provider: "java",
			},
			{
				Name:     "org.jboss.logging.jboss-logging",
				Version:  "3.3.2.Final",
				Provider: "java",
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "JSF", Category: api.Ref{Name: "MVC"}},
		{Name: "Properties", Category: api.Ref{Name: "Other"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "JSF Page", Category: api.Ref{Name: "Web"}},
		{Name: "JSF", Category: api.Ref{Name: "Embedded"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Web"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF Page", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "JSF XML", Category: api.Ref{Name: "View"}},
		{Name: "JSF", Category: api.Ref{Name: "View"}},
		{Name: "JSF Page", Category: api.Ref{Name: "View"}},
		{Name: "Properties", Category: api.Ref{Name: "Embedded"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Properties", Category: api.Ref{Name: "Sustain"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JSF XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB", Category: api.Ref{Name: "Sustain"}},
		{Name: "EJB", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB", Category: api.Ref{Name: "Clustering"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Execute"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Java EE XML", Category: api.Ref{Name: "Processing"}},
	},
}
