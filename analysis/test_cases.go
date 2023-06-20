package analysis

import (
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/api/application"
)

//
// Test application analysis
var BasicTests = []TC{
	{
		Name:        "Pathfinder cloud-readiness",
		Application: application.PathfinderGit,
		Task: api.Task{
			Addon: "windup",
			State: "Ready",
		},
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
		Task: api.Task{
			Addon: "windup",
			State: "Ready",
		},
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
