package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/data/identity"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var Coolstore = TC{
	Name:        "Coolstore",
	Application: data.Coolstore,
	Task:        Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=jakarta-ee",
		},
	},
	Identities: []api.Identity{
		identity.TackleTestappPublicMaven,
	},
	Analysis: api.Analysis{
		Effort: 113,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "HTTP session replication (distributable web.xml)",
				Effort:      3,
				RuleSet:     "cloud-readiness",
				Rule:        "session-00000",
				Incidents: []api.Incident{
					{
						File:     "/shared/source/coolstore/src/main/webapp/WEB-INF/web.xml",
						Line:     5,
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "javax groupId has been replaced by jakarta.platform",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-dependencies-00006",
				Incidents: []api.Incident{
					{
						File:     "/shared/source/coolstore/pom.xml",
						Line:     19,
						Message:  "Update group dependency by replacing the javax groupId with jakarta.platform",
					},
					{
						File:     "/shared/source/coolstore/pom.xml",
						Line:     25,
						Message:  "Update group dependency by replacing the javax groupId with jakarta.platform",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "javax javaee-api artifactId has been replaced by jakarta.platform jakarta.jakartaee-api",
				Effort:      1,
				RuleSet:     "javax-to-jakarta-dependencies-00007",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/pom.xml",
						Line:    26,
						Message: "Update artifact dependency by replacing the javaee-api artifactId with jakarta.jakartaee-api",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "javax javaee-web-api artifactId has been replaced by jakarta.platform jakarta.jakartaee-api",
				Effort:      1,
				RuleSet:     "javax-to-jakarta-dependencies-00008",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/pom.xml",
						Line:    20,
						Message: "Update artifact dependency by replacing the javaee-web-api artifactId with jakarta.jakartaee-web-api",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Rename properties prefixed by javax with jakarta",
				Effort:      1,
				RuleSet:     "javax-to-jakarta-properties-00001",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    10,
						Message: "Rename properties prefixed by javax with jakarta",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE namespace, schemaLocation and version with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "javaee-to-jakarta-namespaces-00001",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    18,
						Message: "Replace http://xmlns.jcp.org/xml/ns/javaee with https://jakarta.ee/xml/ns/jakartaee and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    20,
						Message: "Replace http://xmlns.jcp.org/xml/ns/javaee with https://jakarta.ee/xml/ns/jakartaee and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    21,
						Message: "Replace http://xmlns.jcp.org/xml/ns/javaee with https://jakarta.ee/xml/ns/jakartaee and change the schema version number",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE persistence namespace, schemaLocation and version with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "javaee-to-jakarta-namespaces-00002",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    3,
						Message: "Replace http://xmlns.jcp.org/xml/ns/persistence with https://jakarta.ee/xml/ns/persistence and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    5,
						Message: "Replace http://xmlns.jcp.org/xml/ns/persistence with https://jakarta.ee/xml/ns/persistence and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    6,
						Message: "Replace http://xmlns.jcp.org/xml/ns/persistence with https://jakarta.ee/xml/ns/persistence and change the schema version number",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE version with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "javaee-to-jakarta-namespaces-00033",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    2,
						Message: "In the root tag, replace the version attribute value 2.1 with 3.0",
					},
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    8,
						Message: "In the root tag, replace the version attribute value 2.1 with 3.0",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE XSD with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "javaee-to-jakarta-namespaces-00006",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    21,
						Message: "Replace beans_1_1.xsd with beans_3_0.xsd and update the version attribute to \"3.0\"",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE XSD with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "javaee-to-jakarta-namespaces-00030",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    6,
						Message: "Replace persistence_2_1.xsd with persistence_3_0.xsd and update the version attribute to \"3.0\"",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "The package 'javax' has been replaced by 'jakarta'.",
				Effort:      1,
				RuleSet:     "javax-to-jakarta-import-00001",
				Rule:        "eap8/eap7",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    5,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    6,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    7,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    8,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    9,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    10,
						Message: "Replace the javax.persistence import statement with jakarta.xml",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    7,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    8,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    9,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    10,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    11,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    12,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    13,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    14,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    15,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    5,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    6,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    7,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    8,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    9,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/ShoppingCart.java",
						Line:    7,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/persistence/Resources.java",
						Line:    3,
						Message: "Replace the javax.enterprise import statement with jakarta.enterprise",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/persistence/Resources.java",
						Line:    4,
						Message: "Replace the javax.enterprise import statement with jakarta.enterprise",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/persistence/Resources.java",
						Line:    5,
						Message: "Replace the javax.enterprise import statement with jakarta.enterprise",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/ShoppingCart.java",
						Line:    6,
						Message: "Replace the javax.persistence import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    9,
						Message: "Replace the javax.enterprise import statement with jakarta.persistence",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    10,
						Message: "Replace the javax.inject import statement with jakarta.inject",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    11,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    12,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    13,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    14,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    15,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    16,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    17,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    6,
						Message: "Replace the javax.enterprise import statement with jakarta.enterprise",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    7,
						Message: "Replace the javax.inject import statement with jakarta.inject",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    8,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    9,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    10,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    11,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    12,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    13,
						Message: "Replace the javax.ws import statement with jakarta.ws",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/ProductEndpoint.java",
						Line:    6,
						Message: "Replace the javax.enterprise import statement with jakarta.enterprise",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/ProductEndpoint.java",
						Line:    7,
						Message: "Replace the javax.inject import statement with jakarta.inject",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/ProductEndpoint.java",
						Line:   9,
						Message: "Replace the javax.inject import statement with jakarta.inject",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/ProductEndpoint.java",
						Line:   9,
						Message: "Replace the javax.inject import statement with jakarta.inject",
					},
				},
			},
		},
	    Dependencies: []api.TechDependency{
			{
				Name:     "org.jboss.spec.javax.rmi.jboss-rmi-api_1.0_spec",
				SHA:      "e87e6dae34a2823c6e2e5a0682d3e6aec8abdaa4",
				Version:  "1.0.2.Final",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.jboss.spec.javax.jms.jboss-jms-api_2.0_spec",
				Version:  "2.0.0.Final",
				SHA:      "a67a2cbb27757989592004022f7356e796989340",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "org.flywaydb.flyway-core",
				Version:  "4.1.2",
				SHA:      "e8111260b0111a7a281767829f0aa0a9b27487d4",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "javax.javaee-web-api",
				Version:  "7.0",
				Indirect: true,
				SHA:      "104d7dd93a11b90f7da5e277db912e56460baa32",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "javax.javaee-api",
				Version:  "7.0",
				SHA:      "e54c73cdc03722ed6c66d02cf27d10970ca0cec1",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "com.sun.mail.javax.mail",
				Version:  "1.5.0",
				SHA:      "2143163e070fcd301fbf7b3246ca4854caaabacc",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
			{
				Name:     "javax.activation.activation",
				Version:  "1.1",
				Indirect: true,
				SHA:      "fd9dd0faa8f03f3ce0dc4eec22e57e818d8b9897",
				Provider: "java",
				Labels:   []string{"konveyor.io/dep-source=open-source", "konveyor.io/language=java"},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "EJB XML", Category: api.Ref{Name: "Bean"}},
		{Name: "Web Session", Category: api.Ref{Name: "Clustering"}},
		{Name: "CDI", Category: api.Ref{Name: "Inversion of Control"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Persistence"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Persistence"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Persistence"}},
		{Name: "RMI", Category: api.Ref{Name: "Other"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Processing"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Store"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Store"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Store"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Store"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Connect"}},
		{Name: "RMI", Category: api.Ref{Name: "Connect"}},
		{Name: "Apache License 2.0", Category: api.Ref{Name: "License"}},
		{Name: "GNU GPL", Category: api.Ref{Name: "License"}},
		{Name: "RMI", Category: api.Ref{Name: "Java EE"}},
		{Name: "EJB XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA XML", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA named queries", Category: api.Ref{Name: "Java EE"}},
		{Name: "Persistence units", Category: api.Ref{Name: "Java EE"}},
		{Name: "JPA entities", Category: api.Ref{Name: "Java EE"}},
		{Name: "Web Session", Category: api.Ref{Name: "Java EE"}},
		{Name: "CDI", Category: api.Ref{Name: "Java EE"}},
		{Name: "Web Session", Category: api.Ref{Name: "Sustain"}},
		{Name: "CDI", Category: api.Ref{Name: "Execute"}},
		{Name: "Java EE JSON-P", Category: api.Ref{Name: "Execute"}},
	},
}
