package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var Analyze = api.Task{
	State: "Created", // Created / Ready
	Addon: "analyzer",
}

var AnalyzeDataDefault = addon.Data{
	Output: "/windup/report",
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
				"cloud-readiness",
			},
		},
	},
	Tagger: addon.Tagger{
		Enabled: true,
	},
}
