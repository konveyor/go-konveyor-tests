package utils

import (
	"github.com/jortel/go-utils/logr"
	"github.com/konveyor/tackle2-hub/binding"
	"github.com/konveyor/tackle2-hub/test/api/client"
)

const (
	RETRY_NUM = 10
)

var (
	Log           = logr.WithName("test")
	Client        *binding.Client
	RichClient    *binding.RichClient
	Application   binding.Application
	Tracker       binding.Tracker
	Identity      binding.Identity
	MigrationWave binding.MigrationWave
	Ticket        binding.Ticket
)

func init() {

	// Prepare RichClient and login to Hub API (configured from env variables).
	RichClient = client.PrepareRichClient()

	// Access REST client directly (some test API call need it)
	Client = RichClient.Client

	// Shortcut for Application-related RichClient methods.
	Application = RichClient.Application

	// Shortcut for Tracker-related RichClient methods.
	Tracker = RichClient.Tracker

	Identity = RichClient.Identity
	MigrationWave = RichClient.MigrationWave
	Ticket = RichClient.Ticket
}
