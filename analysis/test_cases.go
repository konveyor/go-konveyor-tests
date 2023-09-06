package analysis

//
// Tier 0 Analysis test cases - should never fail.
// List of applications with expected analysis outputs.
var Tier0TestCases = []TC{
	Tomcat,
	PathfinderSample,
}

//
// Tier 1 Analysis test cases - should work.
// List of applications with expected analysis outputs.
var Tier1TestCases = []TC{
	PetclinicMain,
	PetclinicHazelcast,
	ApacheWicket,
	SeamBooking,
}

//
// Tier 2 Analysis test cases - great if works.
// List of applications with expected analysis outputs.
var Tier2TestCases = []TC{
	Daytrader,
}

//
// Switch analyzers:
// - AnalyzeLsp
// - AnalyzeWindup
var Analyze = AnalyzeLsp
