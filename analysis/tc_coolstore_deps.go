package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var CoolstoreWithDeps = TC{
	Name:        "Coolstore-Source and dependencies",
	Application: data.Coolstore,
	Task:        Analyze,
	WithDeps:    true,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
			"konveyor.io/target=jakarta-ee",
		},
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
						File: "/shared/source/coolstore/src/main/webapp/WEB-INF/web.xml",
						Line: 5,
						Message: `Session replication ensures that client sessions are not disrupted by node failure. Each node in the cluster shares information about ongoing sessions and can take over sessions if another node disappears. In a cloud environment, however, data in the memory of a running container can be wiped out by a restart.

 Recommendations

 * Review the session replication usage and ensure that it is configured properly.
 * Disable HTTP session clustering and accept its implications.
 * Re-architect the application so that sessions are stored in a cache backing service or a remote data grid.

 A remote data grid has the following benefits:

 * The application is more scaleable and elastic.
 * The application can survive EAP node failures because a JVM failure does not cause session data loss.
 * Session data can be shared by multiple applications.`,
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
						File:    "/shared/source/coolstore/pom.xml",
						Line:    19,
						Message: "Update group dependency by replacing the `javax` groupId with `jakarta.platform`",
					},
					{
						File:    "/shared/source/coolstore/pom.xml",
						Line:    25,
						Message: "Update group dependency by replacing the `javax` groupId with `jakarta.platform`",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "javax javaee-api artifactId has been replaced by jakarta.platform jakarta.jakartaee-api",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-dependencies-00007",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/pom.xml",
						Line:    26,
						Message: "Update artifact dependency by replacing the `javaee-api` artifactId with `jakarta.jakartaee-api`",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "javax javaee-web-api artifactId has been replaced by jakarta.platform jakarta.jakartaee-web-api",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-dependencies-00008",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/pom.xml",
						Line:    20,
						Message: "Update artifact dependency by replacing the `javaee-web-api` artifactId with `jakarta.jakartaee-web-api`",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Rename properties prefixed by javax with jakarta",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-properties-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    10,
						Message: "Rename properties prefixed by `javax` with `jakarta`",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE namespace, schemaLocation and version with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javaee-to-jakarta-namespaces-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    18,
						Message: "Replace `http://xmlns.jcp.org/xml/ns/javaee` with `https://jakarta.ee/xml/ns/jakartaee` and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    20,
						Message: "Replace `http://xmlns.jcp.org/xml/ns/javaee` with `https://jakarta.ee/xml/ns/jakartaee` and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    21,
						Message: "Replace `http://xmlns.jcp.org/xml/ns/javaee` with `https://jakarta.ee/xml/ns/jakartaee` and change the schema version number",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE persistence namespace, schemaLocation and version with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javaee-to-jakarta-namespaces-00002",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    3,
						Message: "Replace `http://xmlns.jcp.org/xml/ns/persistence` with `https://jakarta.ee/xml/ns/persistence` and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    5,
						Message: "Replace `http://xmlns.jcp.org/xml/ns/persistence` with `https://jakarta.ee/xml/ns/persistence` and change the schema version number",
					},
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    6,
						Message: "Replace `http://xmlns.jcp.org/xml/ns/persistence` with `https://jakarta.ee/xml/ns/persistence` and change the schema version number",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE version with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javaee-to-jakarta-namespaces-00033",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    2,
						Message: "In the root tag, replace the `version` attribute value `2.1` with `3.0`",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE XSD with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javaee-to-jakarta-namespaces-00006",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/webapp/WEB-INF/beans.xml",
						Line:    21,
						Message: "Replace `beans_1_1.xsd` with `beans_3_0.xsd` and update the version attribute to `\"3.0\"`",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "Replace the Java EE XSD with the Jakarta equivalent",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javaee-to-jakarta-namespaces-00030",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/resources/META-INF/persistence.xml",
						Line:    6,
						Message: "Replace `persistence_2_1.xsd` with `persistence_3_0.xsd` and update the version attribute to `\"3.0\"`",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "The package 'javax' has been replaced by 'jakarta'.",
				Effort:      1,
				RuleSet:     "eap8/eap7",
				Rule:        "javax-to-jakarta-import-00001",
				Incidents: []api.Incident{
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    5,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    6,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    7,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    8,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    9,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/InventoryEntity.java",
						Line:    10,
						Message: "Replace the `javax.xml` import statement with `jakarta.xml`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    7,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    8,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    9,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    10,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    11,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    12,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    13,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    14,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/Order.java",
						Line:    15,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    5,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    6,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    7,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    8,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/OrderItem.java",
						Line:    9,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/model/ShoppingCart.java",
						Line:    7,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/persistence/Resources.java",
						Line:    3,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/persistence/Resources.java",
						Line:    4,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/persistence/Resources.java",
						Line:    5,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/persistence/Resources.java",
						Line:    6,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    9,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    10,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    11,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    12,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    13,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    14,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    15,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    16,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/CartEndpoint.java",
						Line:    17,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    6,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    7,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    8,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    9,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    10,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    11,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    12,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/OrderEndpoint.java",
						Line:    13,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/ProductEndpoint.java",
						Line:    6,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/ProductEndpoint.java",
						Line:    7,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/ProductEndpoint.java",
						Line:    9,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/RestApplication.java",
						Line:    3,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/rest/RestApplication.java",
						Line:    4,
						Message: "Replace the `javax.ws` import statement with `jakarta.ws`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/CatalogService.java",
						Line:    12,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/CatalogService.java",
						Line:    6,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/CatalogService.java",
						Line:    8,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/CatalogService.java",
						Line:    9,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/CatalogService.java",
						Line:    10,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/CatalogService.java",
						Line:    13,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/InventoryNotificationMDB.java",
						Line:    6,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderService.java",
						Line:    5,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderService.java",
						Line:    6,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderService.java",
						Line:    7,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderService.java",
						Line:    8,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderService.java",
						Line:    9,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderService.java",
						Line:    10,
						Message: "Replace the `javax.persistence` import statement with `jakarta.persistence`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderServiceMDB.java",
						Line:    3,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderServiceMDB.java",
						Line:    4,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderServiceMDB.java",
						Line:    5,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderServiceMDB.java",
						Line:    6,
						Message: "Replace the `javax.jms` import statement with `jakarta.jms`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderServiceMDB.java",
						Line:    7,
						Message: "Replace the `javax.jms` import statement with `jakarta.jms`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderServiceMDB.java",
						Line:    8,
						Message: "Replace the `javax.jms` import statement with `jakarta.jms`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/OrderServiceMDB.java",
						Line:    9,
						Message: "Replace the `javax.jms` import statement with `jakarta.jms`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ProductService.java",
						Line:    7,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ProductService.java",
						Line:    8,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/PromoService.java",
						Line:    9,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShippingService.java",
						Line:    6,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShippingService.java",
						Line:    7,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShoppingCartOrderProcessor.java",
						Line:    5,
						Message: "Replace the `javax.annotation` import statement with `jakarta.annotation`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShoppingCartOrderProcessor.java",
						Line:    4,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShoppingCartOrderProcessor.java",
						Line:    6,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShoppingCartOrderProcessor.java",
						Line:    7,
						Message: "Replace the `javax.jms` import statement with `jakarta.jms`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShoppingCartOrderProcessor.java",
						Line:    8,
						Message: "Replace the `javax.jms` import statement with `jakarta.jms`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShoppingCartService.java",
						Line:    6,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/service/ShoppingCartService.java",
						Line:    7,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/DataBaseMigrationStartup.java",
						Line:    6,
						Message: "Replace the `javax.annotation` import statement with `jakarta.annotation`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/DataBaseMigrationStartup.java",
						Line:    7,
						Message: "Replace the `javax.annotation` import statement with `jakarta.annotation`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/DataBaseMigrationStartup.java",
						Line:    8,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/DataBaseMigrationStartup.java",
						Line:    9,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/DataBaseMigrationStartup.java",
						Line:    10,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/DataBaseMigrationStartup.java",
						Line:    11,
						Message: "Replace the `javax.ejb` import statement with `jakarta.ejb`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/DataBaseMigrationStartup.java",
						Line:    12,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Producers.java",
						Line:    3,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Producers.java",
						Line:    4,
						Message: "Replace the `javax.enterprise` import statement with `jakarta.enterprise`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/StartupListener.java",
						Line:    6,
						Message: "Replace the `javax.inject` import statement with `jakarta.inject`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Transformers.java",
						Line:    12,
						Message: "Replace the `javax.json` import statement with `jakarta.json`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Transformers.java",
						Line:    13,
						Message: "Replace the `javax.json` import statement with `jakarta.json`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Transformers.java",
						Line:    14,
						Message: "Replace the `javax.json` import statement with `jakarta.json`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Transformers.java",
						Line:    15,
						Message: "Replace the `javax.json` import statement with `jakarta.json`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Transformers.java",
						Line:    16,
						Message: "Replace the `javax.json` import statement with `jakarta.json`",
					},
					{
						File:    "/shared/source/coolstore/src/main/java/com/redhat/coolstore/utils/Transformers.java",
						Line:    17,
						Message: "Replace the `javax.json` import statement with `jakarta.json`",
					},
				},
			},
		},
		Dependencies: []api.TechDependency{
			{
				Name:     "org.jboss.spec.javax.rmi.jboss-rmi-api_1.0_spec",
				Version:  "1.0.2.Final",
				Provider: "java",
			},
			{
				Name:     "org.jboss.spec.javax.jms.jboss-jms-api_2.0_spec",
				Version:  "2.0.0.Final",
				Provider: "java",
			},
			{
				Name:     "org.flywaydb.flyway-core",
				Version:  "4.1.2",
				Provider: "java",
			},
			{
				Name:     "javax.javaee-web-api",
				Version:  "7.0",
				Provider: "java",
			},
			{
				Name:     "javax.javaee-api",
				Version:  "7.0",
				Provider: "java",
			},
			{
				Name:     "com.sun.mail.javax.mail",
				Version:  "1.5.0",
				Provider: "java",
			},
			{
				Name:     "javax.activation.activation",
				Version:  "1.1",
				Provider: "java",
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
		{Name: "Common Annotations", Category: api.Ref{Name: "Java EE"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Connect"}},
		{Name: "Common Annotations", Category: api.Ref{Name: "Other"}},
	},
}
