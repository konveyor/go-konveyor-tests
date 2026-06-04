package analysiswindup

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/konveyor/go-konveyor-tests/analysis"
	"github.com/konveyor/go-konveyor-tests/hack/addonwindup"
	"github.com/konveyor/tackle2-hub/shared/api"
	"github.com/konveyor/go-konveyor-tests/utils/testutil"
)

var AnalyzeWindup = api.Task{
	State: "Ready", // Created / Ready
	Data:  defaultWindupData,
	Addon: "windup",
}

var defaultWindupData = addonwindup.Data{
	Output: "/windup/report",
	Mode: addonwindup.Mode{
		Artifact: "",
		Binary:   false,
		WithDeps: false,
		Diva:     true,
	},
	Sources: []string{},
	Targets: []string{},
	Scope: addonwindup.Scope{
		WithKnown: false,
		//Packages: {
		//	Included: []string{},
		//	Excluded: []string{},
		//},
	},
	Rules: addonwindup.Rules{
		Path: "",
		Labels: []string{
			"cloud-readiness",
			"linux",
		},
		//Tags: {
		//	Excluded: []string{},
		//},
	},
	Tagger: addonwindup.Tagger{
		Enabled: true,
	},
}

func GetReportText(t *testing.T, tc *analysis.TC, path string) (text string) {
	// Get report file.
	dirName, err := os.MkdirTemp("/tmp", tc.Name)
	testutil.Must(t, err)
	fileName := filepath.Join(dirName, filepath.Base(path))
	err = RichClient.Application.Bucket(tc.Application.ID).Get(path, dirName)
	testutil.Must(t, err)
	content, err := os.ReadFile(fileName)
	testutil.Must(t, err)

	// Prepare content - strip tags etc.
	tags := regexp.MustCompile(`<.*?>`)
	spaces := regexp.MustCompile(`(\t|  +|\n\t+\n)`)
	text = tags.ReplaceAllString(string(content), "")
	text = spaces.ReplaceAllString(text, "")
	return
}

func CheckWindupReportContent(t *testing.T, tc *analysis.TC) {
	// Check the report content.
	for path, expectedElems := range tc.ReportContent {
		content := GetReportText(t, tc, path)
		// Check its content.
		for _, expectedContent := range expectedElems {
			if !strings.Contains(content, expectedContent) {
				t.Errorf("Error report contect check for %s. Cannot find %s in %s", path, expectedContent, content)
			}
		}
	}
}
