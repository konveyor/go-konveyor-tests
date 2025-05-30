package analysis

// Tier 0 Analysis test cases - should never fail.
// List of applications with expected analysis outputs.
var Tier0TestCases = []TC{
	TackleTestappPublicWithDeps,
	TackleTestappPublicPackageFilter,
	Tomcat,
	CoolstoreWithDeps,
	DaytraderWithDeps,
	AcmeairWebappBinary,          // Binary upload
	TackleTestappPublicBinary,    // Binary fetched from tackle-testapp-public maven registry
}

// Tier 1 Analysis test cases - should work.
// List of applications with expected analysis outputs.
var Tier1TestCases = []TC{
	CoolstoreWithDepsQuarkus,
	AdministracionEfectivoBinary, // Binary upload
}

// Tier 2 Analysis test cases - great if works.
// List of applications with expected analysis outputs.
var Tier2TestCases = []TC{
	PetclinicHazelcast,
	ApacheWicket,
	SeamBooking,
}

// Tier 3 Analysis with credentials test cases - should work
// List of applications with expected analysis outputs.
var Tier3TestCases = []TC{}
