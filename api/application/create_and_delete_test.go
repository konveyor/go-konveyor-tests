package main

import (
	"github.com/konveyor/go-konveyor-tests/pkg/application"
	"github.com/konveyor/go-konveyor-tests/pkg/auth"
	"github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
	"testing"
)

func TestCreateAndDeleteApplication(t *testing.T) {
	type test struct {
		name        string
		application api.Application
		shouldError bool
	}
	tests := []test{
		{
			name: "create valid book server application",
			application: api.Application{
				Name: "book-server",
				Repository: &api.Repository{
					Kind: "git",
					URL:  "https://github.com/ibraginsky/book-server",
				},
			},
			shouldError: false,
		},
		{
			name: "create invalid application",
			application: api.Application{
				Name: "bad-app",
				Repository: &api.Repository{
					Kind: "foobar",
					URL:  "https://foo.com/foobar",
				},
			},
			shouldError: true,
		},
		{
			name: "create valid private application",
			application: api.Application{
				Name:       "tackle-testapp",
				Identities: []api.Ref{},
				Repository: &api.Repository{
					Kind: "git",
					URL:  "https://github.com/konveyor/tackle-testapp",
				},
			},
			shouldError: false,
		},
	}
	client := addon.NewClient("http://192.168.49.2/hub/", "")
	token, err := auth.Login(client, "admin", "foobar")
	if err != nil {
		t.Fatalf("unable to authenticate to hub: %v", err.Error())
	}
	client.SetToken(token)
	for _, tc := range tests {
		// POST /applications
		// Verify success (optionally check GET /applications to confirm in DB)
		// Delete application
		t.Log(tc.name) // Placeholder
		err = application.CreateApplication(client, &tc.application)
		if err != nil && !tc.shouldError {
			t.Fatalf("unexpected error: %v", err.Error())
		}
		if err == nil && tc.shouldError {
			t.Fatalf("expected error and didn't get it")
		}
		err = application.DeleteApplication(client, tc.application.ID)
		if err != nil && !tc.shouldError {
			t.Fatalf("unexpected error: %v", err.Error())
		}
	}
}
