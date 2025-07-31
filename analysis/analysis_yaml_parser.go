package analysis

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/konveyor/tackle2-hub/api"
)

// Application represents the application configuration section in tc.yaml files.
type Application struct {
	Name       string `yaml:"name"`
	Repository struct {
		URL  string `yaml:"url"`
		Kind string `yaml:"kind"`
	} `yaml:"repository"`
}

// TCYamlData represents the structure of tc.yaml files for parsing
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

// PopulateTCFromYaml loads configuration from a tc.yaml file and populates the corresponding
// fields in a TC (Test Case) struct. This enables a hybrid approach where test cases can be
// partially defined in Go code and completed with YAML configuration.
//
// The function looks for tc.yaml files in the CI shared tests directory structure:
// {ciRepoDir}/shared_tests/{tc.Name}/tc.yaml
//
// Fields are only populated if they are empty/unset in the TC struct, allowing for selective override behavior.
//
// Returns an error if:
// - The tc.yaml file cannot be read or parsed (missing files are ignored)
func PopulateTCFromYaml(tc *TC, ciRepoDir string) error {
	// Construct the path to the tc.yaml file
	tcPath := filepath.Join(ciRepoDir, "shared_tests", tc.Name, "tc.yaml")

	// Check if file exists
	if _, err := os.Stat(tcPath); os.IsNotExist(err) {
		return nil
	}

	// Load and parse the YAML file
	var yamlData TCYamlData
	err := loadYAMLFromFile(tcPath, &yamlData)
	if err != nil {
		return fmt.Errorf("failed to parse tc.yaml for test case '%s': %w", tc.Name, err)
	}

	// Update the Application fields
	if tc.Application.Name == "" && tc.Application.Repository == nil {
		tc.Application.Name = yamlData.Application.Name
		tc.Application.Repository = &api.Repository{
			URL:  yamlData.Application.Repository.URL,
			Kind: yamlData.Application.Repository.Kind,
		}
	}
	// Convert sources and targets from YAML to labels
	if len(tc.Labels.Included) == 0 {
		var labels []string

		// Convert sources to konveyor.io/source=<value> labels
		for _, source := range yamlData.Sources {
			labels = append(labels, fmt.Sprintf("konveyor.io/source=%s", source))
		}

		// Convert targets to konveyor.io/target=<value> labels
		for _, target := range yamlData.Targets {
			labels = append(labels, fmt.Sprintf("konveyor.io/target=%s", target))
		}

		// Set the labels in the TC struct
		tc.Labels.Included = labels
	}

	// Set the withDeps flag
	if !tc.WithDeps {
		tc.WithDeps = yamlData.WithDeps
	}

	return nil
}
