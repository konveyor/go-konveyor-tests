package analysis

//
// Tier 0 Analysis test cases - should never fail.
// List of applications with expected analysis outputs.
var Tier0TestCases = []TC{
	TackleTestappPublic,
	TackleTestappPublicCorpFrameworkConfig,
}

//
// Tier 1 Analysis test cases - should work.
// List of applications with expected analysis outputs.
var Tier1TestCases = []TC{
	// Setting empty until have working applications.
}

//
// Tier 2 Analysis test cases - great if works.
// List of applications with expected analysis outputs.
var Tier2TestCases = []TC{
	TackleTestappPublicWithDeps,
	TackleTestappPublicCorpFrameworkConfig,	// Move to tier0 or tier1 once confirmed with Ramon&Pranav
	Tomcat,
	Daytrader,
	PetclinicHazelcast,
	ApacheWicket,
	SeamBooking,
}
