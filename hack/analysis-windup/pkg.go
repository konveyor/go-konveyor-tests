package analysiswindup

import (
	"time"

	"github.com/konveyor/tackle2-hub/shared/binding"
	"github.com/konveyor/go-konveyor-tests/utils/testutil"
)

var (
	// Setup Hub API client
	Client     binding.RestClient
	RichClient *binding.RichClient

	// Analysis waiting loop 5 minutes (60 * 5s)
	Retry = 100
	Wait  = 5 * time.Second
)

func init() {
	// Prepare RichClient and login to Hub API (configured from env variables).
	RichClient = testutil.PrepareRichClient()

	// Access REST client directly (some test API call need it)
	Client = RichClient.Client
}

