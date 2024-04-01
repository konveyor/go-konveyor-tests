package analysis

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/k0kubun/pp"
	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/go-konveyor-tests/hack/uniq"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/binding"
	"github.com/konveyor/tackle2-hub/test/assert"
)

// Test application analysis
func TestApplicationAnalysis(t *testing.T) {
	// Find right test cases for given Tier.
	testCases := Tier0TestCases
	_, tier1 := os.LookupEnv("TIER1")
	if tier1 {
		testCases = Tier1TestCases
	}
	_, tier2 := os.LookupEnv("TIER2")
	if tier2 {
		testCases = Tier2TestCases
	}
	_, tier3 := os.LookupEnv("TIER3")
	if tier3 {
		testCases = Tier3TestCases
	}
	// Run test cases.
	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			// Prepare parallel execution if env variable PARALLEL is set.
			tc := testcase
			_, parallel := os.LookupEnv("PARALLEL")
			if parallel {
				t.Parallel()
			}

			// Prepare Identities, e.g. for Maven repo
			for idx := range tc.Identities {
				identity := &tc.Identities[idx]
				uniq.IdentityName(identity)
				if identity.Kind == "maven" {
					mvnUser := os.Getenv("MAVEN_TESTAPP_USER")
					mvnToken := os.Getenv("MAVEN_TESTAPP_TOKEN")
					if mvnToken == "" && mvnUser == "" {
						mvnUser = "konveyor-read-only-bot"
						mvnToken = getDefaultToken()
					}
					identity.Settings = strings.Replace(identity.Settings, "GITHUB_USER", mvnUser, 1)
					identity.Settings = strings.Replace(identity.Settings, "GITHUB_TOKEN", mvnToken, 1)
					t.Logf("using mvn user %s", mvnUser)
				}
				assert.Should(t, RichClient.Identity.Create(identity))
				tc.Application.Identities = append(tc.Application.Identities, api.Ref{ID: identity.ID})
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
			// tc.Task.Addon = analyzerAddon
			tc.Task.Application = &api.Ref{ID: tc.Application.ID}
			taskData := tc.Task.Data.(addon.Data)
			//for _, r := range tc.CustomRules {
			//	taskData.Rules = append(taskData.Rules, api.Ref{ID: r.ID, Name: r.Name})
			//}
			if len(tc.Sources) > 0 {
				taskData.Sources = tc.Sources
			}
			if len(tc.Targets) > 0 {
				taskData.Targets = tc.Targets
			}
			if len(tc.Labels.Included) > 0 || len(tc.Labels.Excluded) > 0 {
				taskData.Rules.Labels = tc.Labels
			}
			if tc.Rules.Path != "" { // TODO: better rules handling
				taskData.Rules = tc.Rules
			}
			if tc.WithDeps == true {
				taskData.Mode.WithDeps = true
			}
			if tc.Binary {
				taskData.Mode.Binary = true
			}
			if tc.Scope != nil {
				taskData.Scope = *tc.Scope
			}

			taskData.Mode.Artifact = tc.Artifact
			tc.Task.Data = taskData
			assert.Should(t, RichClient.Task.Create(&tc.Task))

			if tc.Artifact != "" {
				// Upload binary file into the bucket
				bucketContent := RichClient.Bucket.Content(tc.Task.Bucket.ID)
				bucketContent.Put("data"+tc.Artifact, tc.Artifact)
			}

			tc.Task.State = "Ready"
			assert.Should(t, RichClient.Task.Update(&tc.Task))

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

			if task.State == "Running" {
				t.Error("Timed out running the test. Details:")
				pp.Println(task)
			}

			if task.State != "Succeeded" {
				t.Error("Analyze Task failed. Details:")
				pp.Println(task)
			}

			var gotAppAnalyses []api.Analysis
			var gotAnalysis api.Analysis

			// Get LSP analysis directly from Hub API
			analysisPath := binding.Path(api.AppAnalysesRoot).Inject(binding.Params{api.ID: tc.Application.ID})
			assert.Should(t, Client.Get(analysisPath, &gotAppAnalyses))
			if len(gotAppAnalyses) < 1 {
				t.Fatalf("Analysis result not present in Hub.")
			}
			analysisDetailPath := binding.Path(api.AnalysisRoot).Inject(binding.Params{api.ID: gotAppAnalyses[len(gotAppAnalyses)-1].ID})
			assert.Should(t, Client.Get(analysisDetailPath, &gotAnalysis))

			// Filter out non-mandatory issues, TODO(maufart): quickfix until decide if we test potential issues too
			var mandatoryIssues []api.Issue
			for _, issue := range gotAnalysis.Issues {
				if issue.Category == "mandatory" {
					mandatoryIssues = append(mandatoryIssues, issue)
				}
			}
			gotAnalysis.Issues = mandatoryIssues

			_, debug := os.LookupEnv("DEBUG")
			if debug {
				DumpAnalysis(t, tc, gotAnalysis)
			}

			// Check the analysis result (effort, issues, etc).
			if gotAnalysis.Effort != tc.Analysis.Effort {
				t.Errorf("Different effort error. Got %d, expected %d", gotAnalysis.Effort, tc.Analysis.Effort)
			}

			// Ensure stable order of Issues.
			sort.SliceStable(gotAnalysis.Issues, func(a, b int) bool { return gotAnalysis.Issues[a].Rule < gotAnalysis.Issues[b].Rule })
			sort.SliceStable(tc.Analysis.Issues, func(a, b int) bool { return tc.Analysis.Issues[a].Rule < tc.Analysis.Issues[b].Rule })

			// Check the analysis issues
			if len(gotAnalysis.Issues) != len(tc.Analysis.Issues) {
				t.Errorf("Different amount of issues error. Got %d, expected %d.", len(gotAnalysis.Issues), len(tc.Analysis.Issues))
				t.Error("Got:")
				pp.Println(gotAnalysis.Issues)
				t.Error("Expected:")
				pp.Println(tc.Analysis.Issues)
			} else {
				for i, got := range gotAnalysis.Issues {
					expected := tc.Analysis.Issues[i]
					if got.Category != expected.Category || got.RuleSet != expected.RuleSet || got.Rule != expected.Rule || got.Effort != expected.Effort || !strings.HasPrefix(got.Description, expected.Description) {
						t.Errorf("\nDifferent issue error. Got %+v\nExpected %+v.\n\n", got, expected)
					}

					// Incidents check.
					if len(expected.Incidents) == 0 {
						t.Log("Skipping empty expected Incidents check.")
						break
					}
					if len(got.Incidents) != len(expected.Incidents) {
						t.Errorf("Different amount of incidents error. Got %d, expected %d.", len(got.Incidents), len(expected.Incidents))
						t.Error("Got:")
						pp.Println(got.Incidents)
						t.Error("Expected:")
						pp.Println(expected.Incidents)

					} else {
						// Ensure stable order of Incidents.
						sort.SliceStable(got.Incidents, func(a, b int) bool { return got.Incidents[a].File < got.Incidents[b].File })
						sort.SliceStable(expected.Incidents, func(a, b int) bool { return expected.Incidents[a].File < expected.Incidents[b].File })
						for j, gotInc := range got.Incidents {
							expectedInc := expected.Incidents[j]
							if gotInc.File != expectedInc.File {
								t.Errorf("\nDifferent incident.File error. Got %+v\nExpected %+v.\n\n", gotInc.File, expectedInc.File)
							}
							if gotInc.Line != expectedInc.Line {
								t.Errorf("\nDifferent incident.Line error. Got %+v\nExpected %+v.\n\n", gotInc.Line, expectedInc.Line)
							}
							if !strings.HasPrefix(gotInc.Message, expectedInc.Message) {
								t.Errorf("\nDifferent incident.Message error. Got %+v\nExpected %+v.\n\n", gotInc.Message, expectedInc.Message)
							}
						}
					}
				}
			}

			// Ensure stable order of Dependencies.
			sort.SliceStable(gotAnalysis.Dependencies, func(a, b int) bool { return gotAnalysis.Dependencies[a].Name < gotAnalysis.Dependencies[b].Name })
			sort.SliceStable(tc.Analysis.Dependencies, func(a, b int) bool { return tc.Analysis.Dependencies[a].Name < tc.Analysis.Dependencies[b].Name })

			// Check the dependencies (if specified by TestCase).
			if len(tc.Analysis.Dependencies) > 0 {
				if len(gotAnalysis.Dependencies) != len(tc.Analysis.Dependencies) {
					t.Errorf("Different amount of dependencies error. Got %d, expected %d.", len(gotAnalysis.Dependencies), len(tc.Analysis.Dependencies))
					t.Error("Got:")
					pp.Println(gotAnalysis.Dependencies)
					t.Error("Expected:")
					pp.Println(tc.Analysis.Dependencies)
				} else {
					for i, got := range gotAnalysis.Dependencies {
						expected := tc.Analysis.Dependencies[i]
						if got.Name != expected.Name || got.Version != expected.Version || got.Provider != expected.Provider {
							t.Errorf("\nDifferent dependency error. Got %+v\nExpected %+v.\n\n", got, expected)
						}
					}
				}
			} else {
				t.Log("Skipping Dependencies check, nothing is expected.")
			}

			// Check analysis-created Tags.
			gotApp, _ := RichClient.Application.Get(tc.Application.ID)

			if debug {
				DumpTags(t, tc, *gotApp)
			}

			// Resolve TagRefs to Tags.
			gotTags := []api.Tag{}
			for _, tagRef := range gotApp.Tags {
				if tagRef.Source == "Analysis" {
					tag, _ := RichClient.Tag.Get(tagRef.ID)
					gotTags = append(gotTags, *tag)
				}
			}

			// Ensure stable order of Tags.
			sort.SliceStable(gotTags, func(a, b int) bool {
				return gotTags[a].Name+gotTags[a].Category.Name < gotTags[b].Name+gotTags[b].Category.Name
			})
			sort.SliceStable(tc.AnalysisTags, func(a, b int) bool {
				return tc.AnalysisTags[a].Name+tc.AnalysisTags[a].Category.Name < tc.AnalysisTags[b].Name+tc.AnalysisTags[b].Category.Name
			})

			// Check Tags (if specified by TestCase).
			if len(tc.AnalysisTags) > 0 {
				if len(tc.AnalysisTags) != len(gotApp.Tags) {
					t.Errorf("Different Tags amount error. Got: %d, expected: %d.\n", len(gotApp.Tags), len(tc.AnalysisTags))
					notFoundTags, unexpectedTags := getTagsDiff(tc.AnalysisTags, gotTags)
					for _, notFoundTag := range notFoundTags {
						pp.Println("Expected tag not found", notFoundTag)
					}
					for _, unexpectedTag := range unexpectedTags {
						pp.Println("Unexpected tag found\n", unexpectedTag)
					}
				} else {
					for i, got := range gotTags {
						expected := tc.AnalysisTags[i]
						if got.Name != expected.Name || got.Category.Name != expected.Category.Name {
							t.Errorf("\nDifferent tag error. Got %+v\nExpected %+v.\n\n", got, expected)
						}
					}
				}
			} else {
				t.Log("Skipping Tags check, nothing is expected.")
			}

			// Allow skip cleanup to keep applications and analysis results for debugging etc.
			_, keep := os.LookupEnv("KEEP")
			if keep {
				return
			}

			// Cleanup Identities.
			for _, r := range tc.Application.Identities {
				assert.Should(t, RichClient.Identity.Delete(r.ID))
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

func getDefaultToken() string {
	key := sha256.Sum256([]byte("k0nv3y0r.io"))
	enc, _ := hex.DecodeString("4b47536d696c993113c25974e461e2c8ae759dd9ef3c417d6be13e97f57b59b0c39ce49c1613641ec890e84dc7896e31f161747147e7ab3c024f3fbcb645cd8e57eacb6d")
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return ""
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return ""
	}
	return string(decrypted)
}

func getTagsDiff(want, got []api.Tag) (notFound []api.Tag, extras []api.Tag) {
	tagKey := func(t api.Tag) string { return fmt.Sprintf("%s-%s", t.Category.Name, t.Name) }
	wantTags := map[string]api.Tag{}
	gotTags := map[string]api.Tag{}
	for _, gotTag := range got {
		gotTags[tagKey(gotTag)] = gotTag
	}
	for _, wantTag := range want {
		wantTags[tagKey(wantTag)] = wantTag
		// find tags we expect, but not present in output
		if _, found := gotTags[tagKey(wantTag)]; !found {
			notFound = append(notFound, wantTag)
		}
	}
	for _, gotTag := range got {
		// find tags we got, but we didn't expect
		if _, found := wantTags[tagKey(gotTag)]; !found {
			extras = append(extras, gotTag)
		}
	}
	return
}
