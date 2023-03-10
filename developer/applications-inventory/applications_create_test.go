package applicationsinventory

import (
	"fmt"
	"testing"

	"github.com/konveyor/go-konveyor-tests/pkg/util"
	"github.com/konveyor/tackle2-hub/api"
)

func TestApplicationsCreate(t *testing.T) {
	tests := []util.Test{
		{
			Name: "Create sample Pathfinder application",
			Application: api.Application{
				Name:        "Pathfinder",
				Description: "Tackle Pathfinder application.",
				Repository: &api.Repository{
					Kind:   "git",
					URL:    "https://github.com/konveyor/tackle-pathfinder.git",
					Branch: "1.2.0",
				},
			},
			ShouldError: false,
		},
		{
			Name: "Create minimalist application",
			Application: api.Application{
				Name: "App1",
			},
			ShouldError: false,
		},
		{
			Name: "Not create application with empty name",
			Application: api.Application{
				Name:        "",
				Description: "Empty application.",
			},
			ShouldError: true,
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
		if err != nil && !tc.ShouldError {
			t.Errorf("Unexpected application create error: %v", err.Error())
		}
		if err == nil && tc.ShouldError {
			t.Errorf("Expected application create error and didn't get it")
		}
		t.Log(tc.Application)

		// Clean the application
		if err == nil && !tc.ShouldError {
			err = hub.Delete(fmt.Sprintf("%s/%d", api.ApplicationsRoot, tc.Application.ID))
			if err != nil && !tc.ShouldError {
				t.Errorf("Unexpected application delete error: %v", err.Error())
			}
		}
	}
}
