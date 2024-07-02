package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var Analyze = api.Task{
	State: "Created", // Created / Ready
	Kind: "analyzer",
}

var AnalyzeDataDefault = addon.Data{
	Mode: addon.Mode{
		Artifact: "",
		Binary:   false,
		WithDeps: false,
		// Diva:     true,
	},
	Sources: []string{},
	Targets: []string{},
	Scope: addon.Scope{
		WithKnown: false,
	},
	Rules: addon.Rules{
		Path: "",
		Labels: addon.Labels{
			Excluded: []string{},
			Included: []string{
				"konveyor.io/target=cloud-readiness",
			},
		},
	},
	Tagger: addon.Tagger{
		Enabled: true,
	},
}
