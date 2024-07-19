package analysis

// Tier 0 Analysis test cases - should never fail.
// List of applications with expected analysis outputs.
var Tier0TestCases = []TC{
	TackleTestappPublicWithDeps,
	TackleTestappPublicPackageFilter,
	AcmeairWebapp,
	BookServerVerified,
}

// Tier 1 Analysis test cases - should work.
// List of applications with expected analysis outputs.
var Tier1TestCases = []TC{
	Tomcat,
}

// Tier 2 Analysis test cases - great if works.
// List of applications with expected analysis outputs.
var Tier2TestCases = []TC{
	AdministracionEfectivo,
	TackleTestappPublic,
	Daytrader,
	PetclinicHazelcast,
	ApacheWicket,
	SeamBooking,
}

// Tier 3 tests - should work and should be executed on internal/private infrastructure
// List of applications with expected analysis outputs.
var Tier3TestCases = []TC{
	TackleTestappPrivateBinary, // Needs GITHUB_USER and GITHUB_TOKEN env variables with Read access to https://github.com/konveyor/tackle-testapp
}
