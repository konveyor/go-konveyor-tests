package analysis

//
// Analysis test cases.
// List of applications with expected analysis outputs.
var TestCases = []TC{
	PathfinderSample,
	PetclinicMain,
	//PetclinicHazelcast,
	Tomcat,
	Daytrader,
	ApacheWicket,
	SeamBooking,
}

//
// Switch analyzers:
// - AnalyzeLsp
// - AnalyzeWindup
var Analyze = AnalyzeLsp
