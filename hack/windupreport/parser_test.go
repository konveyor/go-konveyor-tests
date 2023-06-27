package windupreport

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/assert"
)

type TC struct {
	Name           string
	ReportRoot     string
	AnalysisResult api.Analysis
	// ReportContent map[string][]string
}

var Samples = []TC{
	{
		Name: "Pathfinder",
		ReportRoot: "windup-report-index.html",
		AnalysisResult: api.Analysis{
			Issues: []api.Issue{
				{
					Name: "fooissue",
				},
			},
			Dependencies: []api.TechDependency{

			},
		},
	},
}

func TestParse(t *testing.T) {
	for _, tc := range Samples {
		f, err := os.Open(tc.ReportRoot)
		if err != nil {
			log.Fatal(err)
		}
		gotAnalysis := Parse(bufio.NewReader(f))
	
		if !assert.FlatEqual(gotAnalysis, tc.AnalysisResult) {
			t.Errorf("Different response error. Got %+v, expected %+v", gotAnalysis, tc.AnalysisResult)
		}
	}
}
