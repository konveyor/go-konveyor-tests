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
	"output": "/windup/report",
	"mode": api.Map{
		"Artifact": "",
		"Binary":   false,
		"WithDeps": false,
		// Diva:     true,
	},
	"scope": api.Map{
		"WithKnown": false,
	},
	"rules": api.Map{
		"path": "",
		"labels": api.Map{
			"excluded": []string{},
			"included": []string{
				"cloud-readiness",
			},
		},
	},
	"tagger": api.Map{
		"enabled": true,
	},
}
