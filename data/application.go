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
			URL:  "https://github.com/konveyor-ecosystem/book-server",
			Branch: "ci-oct2025",	// branch is a Tag name, similar in following apps
		},
	}
	TackleTestappPublic = api.Application{
		Name: "Tackle Testapp public",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor/tackle-testapp-public",
			Branch: "ci-2024",
		},
	}
	CustomerTomcatLegacy = api.Application{
		Name: "Customer Tomcat Legacy",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor/example-applications.git",
			Path: "example-1",
			Branch: "ci-2023",
		},
	}
	Coolstore = api.Application{
		Name: "Coolstore",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor-ecosystem/coolstore",
			Branch: "ci-2024",
		},
	}
	CoolstoreQuarkus = api.Application{
		Name: "Coolstore",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor-ecosystem/coolstore",
			Branch: "ci-2024-quarkus",
		},
	}
	Daytrader = api.Application{
		Name: "Daytrader 7 EE application",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor-ecosystem/sample.daytrader7.git",
		},
	}
	ApacheWicket = api.Application{
		Name: "Apache Wicket",
		Repository: &api.Repository{
			Kind: "git",
			URL:  "https://github.com/konveyor-ecosystem/windup-sample-apps.git",
			Path: "test-files/src_example/org/apache/wicket",
			Branch: "ci-2020",
		},
	}
	PetclinicHazelcast = api.Application{
		Name:        "Petclinic",
		Description: "Spring framework app",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/konveyor/spring-framework-petclinic.git",
			Branch: "ci-2023-legacy",
		},
	}
	SeamBooking = api.Application{
		Name: "Seam booking 5.2",
		Repository: &api.Repository{
			Kind:   "git",
			URL:    "https://github.com/konveyor-ecosystem/windup.git",
			Path:   "test-files/seam-booking-5.2",
			Branch: "ci-2024",
		},
	}

	// Binary applications
	TackleTestappPublicBinary = api.Application{
		Name:   "Tackle Testapp public binary",
		Binary: "mvn://io.konveyor.demo:customers-tomcat:0.0.1-SNAPSHOT:war", // To be downloaded from maven repository
	}
	UploadBinary = api.Application{
		Name: "upload-binary", // The file itself to be uploaded during the test
	}

	ApplicationSamples = []api.Application{Minimal, PathfinderGit, BookServer, TackleTestappPublic,
		CustomerTomcatLegacy, Coolstore, CoolstoreQuarkus, Daytrader, ApacheWicket, PetclinicHazelcast,
		SeamBooking, TackleTestappPublicBinary, UploadBinary}
)
