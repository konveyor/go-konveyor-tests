package analysis

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/go-konveyor-tests/hack/windupreport"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/assert"
)

// Test application analysis
func TestApplicationAnalysis(t *testing.T) {
	for _, analyzerAddon := range Addons {
		for _, testcase := range TestCases {
			t.Run(fmt.Sprintf("%s_%s", testcase.Name, analyzerAddon), func(t *testing.T) {
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
				tc.Task.Addon = analyzerAddon
				tc.Task.Application = &api.Ref{ID: tc.Application.ID}
				taskData := tc.Task.Data.(addon.Data)
				for _, r := range tc.CustomRules {
					taskData.Rules.RuleSets = append(taskData.Rules.RuleSets, api.Ref{ID: r.ID, Name: r.Name})
				}
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

				// Check the report content.
				for path, expectedElems := range tc.ReportContent {
					content := getReportText(t, &tc, path)
					// Check its content.
					for _, expectedContent := range expectedElems {
						if !strings.Contains(content, expectedContent) {
							t.Errorf("Error report contect check for %s. Cannot find %s in %s", path, expectedContent, content)
						}
					}
				}

				// Check the analysis result (effort, issues, etc).
				// Parse report for windup, get analysis from Hub API for lsp analyzer
				gotAnalysis := windupreport.Parse(t, tc.Application.ID)
				if gotAnalysis.Effort != tc.Analysis.Effort {
					t.Errorf("Different effort error. Got %d, expected %d", gotAnalysis.Effort, tc.Analysis.Effort)
				}
				if !assert.FlatEqual(gotAnalysis.Issues, tc.Analysis.Issues) {
					t.Errorf("Analysis Issues don't match. Got:\n  %+v\nexpected:\n  %+v\n", gotAnalysis.Issues, tc.Analysis.Issues)
				}
				for i := range gotAnalysis.Issues {
					if !assert.FlatEqual(gotAnalysis.Issues[i].Incidents, tc.Analysis.Issues[i].Incidents) {
						t.Errorf("Analysis Incidents don't match. Got:\n  %+v\nexpected:\n  %+v\n", gotAnalysis.Issues[i], tc.Analysis.Issues[i])
					}
				}


				// Check analysis-created Tags.
				gotApp, _ := RichClient.Application.Get(tc.Application.ID)
				found, gotAnalysisTags := 0, 0
				for _, t := range gotApp.Tags {
					if t.Source == "Analysis" {
						gotAnalysisTags = gotAnalysisTags + 1
						for _, expectedTag := range tc.AnalysisTags {
							if expectedTag.Name == t.Name {
								found = found + 1
								break
							}
						}
					}
				}
				if found != len(tc.AnalysisTags) || found < gotAnalysisTags {
					t.Errorf("Analysis Tags don't match. Got:\n  %v\nexpected:\n  %v\n", gotApp.Tags, tc.AnalysisTags)
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
}
