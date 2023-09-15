package analysiswindup

import (
	"os"
	"testing"
	"time"

	"github.com/konveyor/go-konveyor-tests/analysis"
	"github.com/konveyor/go-konveyor-tests/hack/addonwindup"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/go-konveyor-tests/hack/windupreport"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/assert"
)

// Test application analysis
func TestApplicationAnalysis(t *testing.T) {
	// Find right test cases for given Tier (include Tier0 always).
	testCases := analysis.Tier0TestCases
	_, tier1 := os.LookupEnv("TIER1")
	if tier1 {
		testCases = analysis.Tier1TestCases
	}
	_, tier2 := os.LookupEnv("TIER2")
	if tier2 {
		testCases = analysis.Tier2TestCases
	}
	// Run test cases.
	for _, testcase := range testCases {
		t.Run("Windup"+testcase.Name, func(t *testing.T) {
			// Prepare parallel execution if env variable PARALLEL is set.
			tc := testcase
			_, parallel := os.LookupEnv("PARALLEL")
			if parallel {
				t.Parallel()
			}

			// Create the application.
			uniq.ApplicationName(&tc.Application)
			assert.Should(t, RichClient.Application.Create(&tc.Application))

			// Prepare custom rules.
			for i := range tc.CustomRules {
				r := &tc.CustomRules[i]
				uniq.RuleSetName(r)
				// ruleFiles := []api.File{}
				rules := []api.Rule{}
				for _, rule := range r.Rules {
					ruleFile, err := RichClient.File.Put(rule.File.Name)
					assert.Should(t, err)
					rules = append(rules, api.Rule{
						File: &api.Ref{
							ID: ruleFile.ID,
						},
					})
					// ruleFiles = append(ruleFiles, *ruleFile)
				}
				r.Rules = rules
				assert.Should(t, RichClient.RuleSet.Create(r))
			}

			// Prepare and submit the analyze task.
			tc.Task = AnalyzeWindup
			tc.Task.Application = &api.Ref{ID: tc.Application.ID}
			taskData := tc.Task.Data.(addonwindup.Data) // addon / addonwindup
			//for _, r := range tc.CustomRules {
			//	taskData.Rules = append(taskData.Rules, api.Ref{ID: r.ID, Name: r.Name})
			//}
			if len(tc.Sources) > 0 {
				taskData.Sources = tc.Sources
			}
			if len(tc.Targets) > 0 {
				taskData.Targets = tc.Targets
			}
			//if len(tc.Labels.Included) > 0 || len(tc.Labels.Excluded) > 0 {
			//	taskData.Rules.Labels = tc.Labels
			//}
			//if tc.Rules.Path != "" { // TODO: better rules handling
			//	taskData.Rules = tc.Rules
			//}
			tc.Task.Data = taskData
			assert.Should(t, RichClient.Task.Create(&tc.Task))

			// Wait until task finishes
			var task *api.Task
			var err error
			for i := 0; i < Retry; i++ {
				task, err = RichClient.Task.Get(tc.Task.ID)
				if err != nil || task.State == "Succeeded" || task.State == "Failed" {
					break
				}
				time.Sleep(Wait)
			}

			if task.State != "Succeeded" {
				t.Errorf("Analyze Task failed. Details: %+v", task)
			}

			var gotAnalysis api.Analysis

			// Old Windup check version parsing windup HTML report
			CheckWindupReportContent(t, &tc)

			// Parse report for windup, to api.Analysis structure and print
			gotAnalysis = windupreport.Parse(t, tc.Application.ID)
			gotApp, _ := RichClient.Application.Get(tc.Application.ID)
			analysis.DumpAnalysis(t, tc, gotAnalysis)
			analysis.DumpTags(t, tc, *gotApp)

			// Allow skip cleanup to keep applications and analysis results for debugging etc.
			_, keep := os.LookupEnv("KEEP")
			if keep {
				return
			}
			// Cleanup Application.
			assert.Must(t, RichClient.Application.Delete(tc.Application.ID))

			// Cleanup custom rules and their files.
			for _, r := range tc.CustomRules {
				assert.Should(t, RichClient.RuleSet.Delete(r.ID))
				for _, rl := range r.Rules {
					assert.Should(t, RichClient.File.Delete(rl.File.ID))
				}
			}
		})
	}
}
