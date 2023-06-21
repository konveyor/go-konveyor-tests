package analysis

import (
	"github.com/konveyor/tackle2-hub/api"
)

//
// Test application analysis
var TestCases = []TC{
	{
		Name:        "Pathfinder cloud-readiness",
		Application: api.Application{
			Name:        "Pathfinder",
			Description: "Tackle Pathfinder application.",
			Repository: &api.Repository{
				Kind:   "git",
				URL:    "https://github.com/konveyor/tackle-pathfinder.git",
				Branch: "1.2.0",
			},
		},
		Task: WindupReady,
		TaskData: defaultTaskData,
		ReportContent: map[string][]string{
			"/windup/report/index.html": {
				"5\nstory points",
				"5\nCloud Mandatory",
				"9\nInformation",
			},
		},
		//ExpectedTags: []api.Tag{
		//	{id: 121, Name: "CDI", source: "Analysis"},
		//},
	},
	{
		Name:        "Petclinic cloud-readiness",
		Application: api.Application{
			Name: "Petclinic",
			Description: "Spring framework app",
			Repository: &api.Repository{
				Kind:   "git",
				URL:    "https://github.com/savitharaghunathan/spring-framework-petclinic.git",
			},
			
		},
		Task: WindupReady,
		TaskData: defaultTaskData,
		ReportContent: map[string][]string{
			"/windup/report/index.html": {
				"5\nstory points",
				"5\nCloud Mandatory",
				"4\nInformation",
			},
		},
	},
}

//
// Shared parameters.
var WindupReady = api.Task{
	Addon: "windup",
	State: "Ready",
}

var defaultTaskData = `{
	"mode": {
		"artifact": "",
		"binary": false,
		"withDeps": false,
		"diva": true
	},
	"output": "/windup/report",
	"rules": {
		"path": "",
		"tags": {
			"excluded": [ ]
		}
	},
	"scope": {
		"packages": {
			"excluded": [ ],
			"included": [ ]
		},
		"withKnown": false
	},
	"sources": [ ],
	"tagger": {
		"enabled": true
	},
	"targets": [
		"cloud-readiness"
	]
  }`
