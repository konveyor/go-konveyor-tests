package analysis

import (
	"fmt"
	"os"

	"github.com/konveyor/tackle2-hub/api"
	"gopkg.in/yaml.v3"
)

// Application represents the application configuration section in test_cases.yml file.
type Application struct {
	Name       string `yaml:"name"`
	Repository struct {
		URL    string `yaml:"url"`
		Kind   string `yaml:"kind"`
		Branch string `yaml:"branch"`
		Tag    string `yaml:"tag"`
		Path   string `yaml:"path"`
	} `yaml:"repository"`
}

// TCYamlData represents the structure of test_cases.yml file for parsing
type TCYamlData struct {
	Description string      `yaml:"description"`
	Application Application `yaml:"application"`
	Input       string      `yaml:"input"`
	Filename    string      `yaml:"filename"`
	Sources     []string    `yaml:"sources"`
	Targets     []string    `yaml:"targets"`
	WithDeps    bool        `yaml:"withDeps"`
}

// AnalysisIncident represents an incident in the analysis output YAML
type AnalysisIncident struct {
	CodeSnip   string `yaml:"codeSnip"`
	LineNumber int    `yaml:"lineNumber"`
	Message    string `yaml:"message"`
	URI        string `yaml:"uri"`
}

// AnalysisViolation represents a violation in the analysis output YAML
type AnalysisViolation struct {
	Category    string             `yaml:"category"`
	Description string             `yaml:"description"`
	Effort      int                `yaml:"effort"`
	Incidents   []AnalysisIncident `yaml:"incidents"`
}

// AnalysisRuleset represents a ruleset in the analysis output YAML
type AnalysisRuleset struct {
	Description string                       `yaml:"description"`
	Name        string                       `yaml:"name"`
	Violations  map[string]AnalysisViolation `yaml:"violations"`
	Tags        []string                     `yaml:"tags"`
}

// OutputYamlData represents the structure of output.yaml file for parsing
type OutputYamlData []AnalysisRuleset

func loadYAMLFromFile(path string, out interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %w", err)
	}
	return yaml.Unmarshal(data, out)
}

// loadTestConfig updates a TC struct with values from the testCasesData map for the given test case name if available	.
// Fields are only populated if they are empty/unset in the TC struct, allowing for selective override behavior.
func loadTestConfig(tc *TC, testCasesData map[string]TCYamlData) error {
	// Get the test case data for the current test case
	testCaseData, ok := testCasesData[tc.Name]
	if !ok {
		return nil
	}

	// Update the Application fields
	if tc.Application.Name == "" && tc.Application.Repository == nil {
		tc.Application.Name = testCaseData.Application.Name
		tc.Application.Repository = &api.Repository{
			URL:    testCaseData.Application.Repository.URL,
			Kind:   testCaseData.Application.Repository.Kind,
			Branch: testCaseData.Application.Repository.Branch,
			Tag:    testCaseData.Application.Repository.Tag,
			Path:   testCaseData.Application.Repository.Path,
		}
	}
	// Convert sources and targets from YAML to labels
	if len(tc.Labels.Included) == 0 {
		var labels []string

		// Convert sources to konveyor.io/source=<value> labels
		for _, source := range testCaseData.Sources {
			labels = append(labels, fmt.Sprintf("konveyor.io/source=%s", source))
		}

		// Convert targets to konveyor.io/target=<value> labels
		for _, target := range testCaseData.Targets {
			labels = append(labels, fmt.Sprintf("konveyor.io/target=%s", target))
		}

		// Set the labels in the TC struct
		tc.Labels.Included = labels
	}

	// Set the withDeps flag
	if !tc.WithDeps {
		tc.WithDeps = testCaseData.WithDeps
	}

	return nil
}

// populateAnalysisOutput updates a TC struct with analysis results from the analysisOutput.
// Fields are only populated if the TC.Analysis is empty/unset, allowing for selective override behavior.
func populateAnalysisOutput(tc *TC, analysisOutput OutputYamlData) error {
	var insights []api.Insight
	var totalEffort int

	// Iterate through each ruleset
	for _, ruleset := range analysisOutput {
		// Iterate through each violation in the ruleset
		for ruleID, violation := range ruleset.Violations {
			insight := api.Insight{
				Category:    violation.Category,
				Description: violation.Description,
				Effort:      violation.Effort,
				RuleSet:     ruleset.Name,
				Rule:        ruleID,
			}

			// Convert YAML incidents to API incidents
			for _, yamlIncident := range violation.Incidents {
				totalEffort += violation.Effort
				incident := api.Incident{
					File:    yamlIncident.URI,
					Line:    yamlIncident.LineNumber,
					Message: yamlIncident.Message,
				}
				insight.Incidents = append(insight.Incidents, incident)
			}
			insights = append(insights, insight)
		}
	}

	// Only populate if empty
	if tc.Analysis.Effort == 0 {
		tc.Analysis.Effort = totalEffort
	}
	if len(tc.Analysis.Insights) == 0 {
		tc.Analysis.Insights = insights
	}
	return nil
}
