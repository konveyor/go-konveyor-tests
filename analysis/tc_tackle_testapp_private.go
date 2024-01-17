package analysis

import (
	"github.com/konveyor/go-konveyor-tests/data"
	"github.com/konveyor/go-konveyor-tests/data/identity"
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPrivateBinary = TC{
	Name:        "tackle-testapp-binary",
	Application: data.TackleTestapp,
	Identities: []api.Identity{
		identity.TackleTestappMaven,
	},
	Task: Analyze,
	Binary: true,
	Analysis: api.Analysis{
		Effort: 16,
		Issues: []api.Issue{
			{Category: "mandatory", Description: "Java native libraries (JNI, JNA)", Effort: 7, Rule: "jni-native-code-00000"},
			{Category: "mandatory", Description: "File system - Java IO", Effort: 1, Rule: "local-storage-00001"},
			{Category: "potential", Description: "Java Servlet", Effort: 0, Rule: "javaee-technology-usage-00120"},
		},
	},
	AnalysisTags: []api.Tag{
		{Name: "Servlet"},
		{Name: "Java Threads"},
	},
}
