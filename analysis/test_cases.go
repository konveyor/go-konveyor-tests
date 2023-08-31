package analysis

//
// Stage 0 Analysis test cases - should never fail.
// List of applications with expected analysis outputs.
var Stage0TestCases = []TC{
	Tomcat,
	//PathfinderSample,
}

//
// Stage 1 Analysis test cases - should work.
// List of applications with expected analysis outputs.
var Stage1TestCases = []TC{
	PetclinicMain,
	PetclinicHazelcast,
	Daytrader,
	ApacheWicket,
	SeamBooking,
}

//
// Stage 2 Analysis test cases - great if works.
// List of applications with expected analysis outputs.
var Stage2TestCases = []TC{
}

//
// Switch analyzers:
// - AnalyzeLsp
// - AnalyzeWindup
var Analyze = AnalyzeLsp
