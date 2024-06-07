package analysis

import (
	"github.com/konveyor/tackle2-hub/api"
)

var Analyze = api.Task{
	State: "Created", // Created / Ready
	Data:  defaultAnalyzerData,
	Addon: "analyzer",
}

var defaultAnalyzerData = api.Map{
	"Output": "/windup/report",
	"Mode": api.Map{
		"Artifact": "",
		"Binary":   false,
		"WithDeps": false,
		// Diva:     true,
	},
	"Scope": api.Map{
		"WithKnown": false,
	},
	"Rules": api.Map{
		"Path": "",
		"Labels": api.Map{
			"Excluded": []string{},
			"Included": []string{
				"cloud-readiness",
			},
		},
	},
	"Tagger": api.Map{
		"Enabled": true,
	},
}
