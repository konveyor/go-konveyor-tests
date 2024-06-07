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
		"artifact": "",
		"binary":   false,
		"withDeps": false,
	},
	"scope": api.Map{
		"withKnown": false,
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
