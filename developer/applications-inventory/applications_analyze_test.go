package applicationsinventory

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/konveyor/go-konveyor-tests/util"
	"github.com/konveyor/tackle2-hub/api"
)

func TestApplicationsAnalyze(t *testing.T) {
	tests := []util.Test{
		{
			Name: "Analyze sample Pathfinder application",
			Application: api.Application{
				Name:        "Pathfinder",
				Description: "Tackle Pathfinder application.",
				Repository: &api.Repository{
					Kind:   "git",
					URL:    "https://github.com/konveyor/tackle-pathfinder.git",
					Branch: "1.2.0",
				},
			},
			Task: api.Task{
				Name:        "Windup test analysis",
				Addon:       "windup",
				State:       "Ready",
				Application: &api.Ref{}, // Application's ID is set in the test execution
				Data:        map[string]string{},
			},
			// Passing Task Data as a string until it is possible import e.g. windup Data type
			TaskData: `{
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
					"targets": [
						"cloud-readiness"
					]
				  }`,
		},
	}

	// Setup Hub API client
	hub, err := util.NewHubClient()
	if err != nil {
		t.Fatalf("Unable connect to Hub API: %v", err.Error())
	}

	// Execute test steps
	for _, tc := range tests {
		t.Log(tc.Name)

		// Create the application
		err = hub.Post(api.ApplicationsRoot, &tc.Application)
		if err != nil {
			t.Errorf("Unexpected application create error: %v", err.Error())
		}
		t.Log(tc.Application)

		// Create task to start the analysis
		tc.Task.Application.ID = tc.Application.ID
		json.Unmarshal([]byte(tc.TaskData), &tc.Task.Data)
		err = hub.Post(api.TasksRoot, &tc.Task)
		if err != nil || tc.Task.State != "Ready" {
			t.Errorf("Unexpected task create error: %v", err.Error())
		}

		// Wait until task finishes, 5 minutes timeout (60*5.seconds)
		for i := 0; i < 60; i++ {
			err = hub.Get(fmt.Sprintf("%s/%d", api.TasksRoot, tc.Task.ID), &tc.Task)
			if err != nil || tc.Task.State == "Succeeded" || tc.Task.State == "Failed" {
				// Proceed to Task result check
				break
			}
			t.Log(tc.Task)
			time.Sleep(5 * time.Second)
		}

		// Check the Task
		err = hub.Get(fmt.Sprintf("%s/%d", api.TasksRoot, tc.Task.ID), &tc.Task)
		if err != nil || tc.Task.State != "Succeeded" {
			t.Errorf("Application analyze Task failed: %v", tc.Task)
		}

		// Check Task's report
		// TODO

		// Clean the application
		err = hub.Delete(fmt.Sprintf("%s/%d", api.ApplicationsRoot, tc.Application.ID))
		if err != nil && !tc.ShouldError {
			t.Errorf("Unexpected application delete error: %v", err.Error())
		}
	}
}
