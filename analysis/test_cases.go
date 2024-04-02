package analysis

// Tier 0 Analysis test cases - should never fail.
// List of applications with expected analysis outputs.
var Tier0TestCases = []TC{
	TackleTestappPublicWithDeps,
	TackleTestappPublicPackageFilter,
}

// Tier 1 Analysis test cases - should work.
// List of applications with expected analysis outputs.
var Tier1TestCases = []TC{
	// Setting empty until have working applications.
}

// Tier 2 Analysis test cases - great if works.
// List of applications with expected analysis outputs.
var Tier2TestCases = []TC{
	AdministracionEfectivo,
	Tomcat,
	TackleTestappPublic,
	Daytrader,
	PetclinicHazelcast,
	ApacheWicket,
	SeamBooking,
}

// Tier 3 Analysis test cases - should work
// List of applications with expected analysis outputs.
var Tier3TestCases = []TC{
	TackleTestappPrivateBinary,
}
