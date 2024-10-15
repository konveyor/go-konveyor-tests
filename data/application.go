package data

import (
	"github.com/konveyor/tackle2-hub/api"
)

// Set of valid Application resources for tests and reuse.
// Important: initialize test application from this samples, not use it directly to not affect other tests.
var (
	Minimal = api.Application{
		Name: "Minimal application",
	}

	// Repository-based applications
	PathfinderGit = api.Application{
		Name:        "Pathfinder",
		Description: "Tackle Pathfinder application.",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/konveyor/tackle-pathfinder.git",
			Branch: "1.2.0",
		},
	}
	BookServer = api.Application{
		Name: "Book Server",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/ibraginsky/book-server",
		},
	}
	TackleTestappPublic = api.Application{
		Name: "Tackle Testapp public",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor/tackle-testapp-public",
		},
	}
	CustomerTomcatLegacy = api.Application{
		Name: "Customer Tomcat Legacy",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor/example-applications.git",
			Path: "example-1",
		},
	}
	Coolstore = api.Application{
		Name: "Coolstore",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor-ecosystem/coolstore",
		},
	}

	// Binary applications
	TackleTestappPublicBinary = api.Application{
		Name: "Tackle Testapp public binary",
		Binary: "mvn://io.konveyor.demo:customers-tomcat:0.0.1-SNAPSHOT:war",	// To be downloaded from maven repository
	}
	UploadBinary = api.Application{
		Name: "upload-binary",	// The file itself to be uploaded during the test
	}

	ApplicationSamples = []api.Application{Minimal, PathfinderGit, BookServer, TackleTestappPublic,
		CustomerTomcatLegacy, Coolstore, TackleTestappPublicBinary, UploadBinary}
)
