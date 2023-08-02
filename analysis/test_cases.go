package analysis

import (
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/go-konveyor-tests/hack/addonwindup"
	"github.com/konveyor/tackle2-hub/api"
)

var TestCases = []TC{
	PathfinderSample,
	PetclinicMain,
	////PetclinicHazelcast,
	Tomcat,
	Daytrader,
	ApacheWicket,
	SeamBooking,
}

var AnalyzeWindup = api.Task{
	State: "Ready", // Created / Ready
	Data:  defaultWindupData,
	Addon: "windup",
}

var AnalyzeLsp = api.Task{
	State: "Ready", // Created / Ready
	Data:  defaultAnalyzerData,
	Addon: "analyzer",
}

//
// Switch analyzers.
var Analyze = AnalyzeLsp

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

var defaultWindupData = addonwindup.Data{
	Output: "/windup/report",
	Mode: addonwindup.Mode{
		Artifact: "",
		Binary:   false,
		WithDeps: false,
		Diva:     true,
	},
	Sources: []string{},
	Targets: []string{"cloud-readiness"},
	Scope: addonwindup.Scope{
		WithKnown: false,
		//Packages: {
		//	Included: []string{},
		//	Excluded: []string{},
		//},
	},
	Rules: addonwindup.Rules{
		Path: "",
		Labels: []string{
			"cloud-readiness",
		},
		//Tags: {
		//	Excluded: []string{},
		//},
	},
	Tagger: addonwindup.Tagger{
		Enabled: true,
	},
}
