package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/data/identity"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPublicBinary = TC{
	Name:        "tackle-testapp-binary",
	Application: data.TackleTestappPublicBinary,
	Identities: []api.Identity{
		identity.TackleTestappPublicMaven,
	},
	Task: Analyze,
	Labels: addon.Labels{
		Included: []string{
			"konveyor.io/target=cloud-readiness",
		},
	},
	Binary: true,
	Analysis: api.Analysis{
		Effort: 16,
		Issues: []api.Issue{
			{
				Category:    "mandatory",
				Description: "Java native libraries (JNI, JNA)",
				Effort:      7,
				RuleSet:     "cloud-readiness",
				Rule:        "jni-native-code-00000",
				Incidents: []api.Incident{
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/config/ApplicationConfiguration.java",
						Line:    17,
						Message: "\n Java native libraries might not run in a cloud or container environment.",
					},
				},
			},
			{
				Category:    "mandatory",
				Description: "File system - Java IO",
				Effort:      1,
				RuleSet:     "cloud-readiness",
				Rule:        "local-storage-00001",
				Incidents: []api.Incident{
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/config/ApplicationConfiguration.java",
						Line:    8,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/config/ApplicationConfiguration.java",
						Line:    37,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    39,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    40,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    41,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    42,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    61,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/ordermanagement/config/PersistenceConfig.java",
						Line:    62,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
					{
						File:    "/addon/bin/maven/java-project/src/main/java/io/konveyor/demo/ordermanagement/exception/handler/ExceptionHandlingController.java",
						Line:    20,
						Message: "\n An application running inside a container could lose access to a file in local storage.",
					},
				},
			},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet", Category: api.Ref{Name: "HTTP"}},
		{Name: "JNI", Category: api.Ref{Name: "Other"}},
		{Name: "Servlet", Category: api.Ref{Name: "Connect"}},
		{Name: "JNI", Category: api.Ref{Name: "Connect"}},
		{Name: "Servlet", Category: api.Ref{Name: "Java EE"}},
		{Name: "JNI", Category: api.Ref{Name: "Java EE"}},
	},
}
