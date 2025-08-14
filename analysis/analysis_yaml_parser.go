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
		URL  string `yaml:"url"`
		Kind string `yaml:"kind"`
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
			URL:  testCaseData.Application.Repository.URL,
			Kind: testCaseData.Application.Repository.Kind,
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
