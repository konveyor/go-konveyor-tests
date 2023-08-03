package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
)

var AnalyzeLsp = api.Task{
	State: "Ready", // Created / Ready
	Data:  defaultAnalyzerData,
	Addon: "analyzer",
}

var defaultAnalyzerData = addon.Data{
	Output: "/windup/report",
	Mode: addon.Mode{
		Artifact: "",
		Binary:   false,
		WithDeps: false,
		//Diva:     true,
	},
	Sources: []string{},
	Targets: []string{"cloud-readiness"},
	Scope: addon.Scope{
		WithKnown: false,
		//Packages: {
		//	Included: []string{},
		//	Excluded: []string{},
		//},
	},
	Rules: addon.Rules{
		Path: "",
		Labels: struct {
			Included []string `json:"included,omitempty"`
			Excluded []string `json:"excluded,omitempty"`
		}{
			Excluded: []string{},
			Included: []string{
			"cloud-readiness",
			},
		},
		//Tags: {
		//	Excluded: []string{},
		//},
	},
	Tagger: addon.Tagger{
		Enabled: true,
	},
}

