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
		Binary: "customers-tomcat-0.0.1-SNAPSHOT.war",
	}
	UploadBinary = api.Application{
		Name: "upload-binary",
	}
	CustomerTomcatLegacy = api.Application{
		Name: "Customer Tomcat Legacy",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor/example-applications.git",
			Path: "example-1",
		},
	}
	ApplicationSamples = []api.Application{Minimal, PathfinderGit, BookServer, TackleTestappPublic, UploadBinary, CustomerTomcatLegacy}
)
