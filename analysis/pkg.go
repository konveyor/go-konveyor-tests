package analysis

import (
	"fmt"
	"testing"
	"time"

	"github.com/konveyor/go-konveyor-tests/hack/addon"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/binding"
	"github.com/konveyor/tackle2-hub/test/api/client"
)

const (
	Username = "HUB_USERNAME"
	Password = "HUB_PASSWORD"
)

var (
	// Setup Hub API client
	Client     *binding.Client
	RichClient *binding.RichClient

	// Analysis waiting loop 5 minutes (60 * 5s)
	Retry = 200
	Wait  = 5 * time.Second
)

func init() {
	// Prepare RichClient and login to Hub API (configured from env variables).
	RichClient = client.PrepareRichClient()

	// Access REST client directly (some test API call need it)
	Client = RichClient.Client
}

// Test cases for Application Analysis.
type TC struct {
	Name string
	// Application and other test data declaration.
	Application api.Application // Required.
	CustomRules []api.RuleSet
	Identities  []api.Identity
	// Analysis parameters.
	Task      api.Task
	TaskData  string
	Sources   []string
	Targets   []string
	Labels    addon.Labels
	RulesPath string
	Scope     *addon.Scope
	WithDeps  bool
	Binary    bool
	Artifact  string
	// After-analysis assertions.
	ReportContent map[string][]string
	Analysis      api.Analysis
	AnalysisTags  []api.Tag
}

func DumpAnalysis(t *testing.T, tc TC, analysis api.Analysis) {
	fmt.Printf("## GOT ANALYSIS OUTPUT FOR \"%s\":", tc.Name)
	fmt.Printf("\napi.Analysis{\n")
	fmt.Printf("    Effort: %d,\n", analysis.Effort)
	fmt.Printf("    Issues: []api.Issue{\n")
	for _, issue := range analysis.Issues {
		fmt.Printf("        {\n")
		fmt.Printf("            Category: \"%s\",\n", issue.Category)
		fmt.Printf("            Description: \"%s\",\n", issue.Description)
		fmt.Printf("            Effort: %d,\n", issue.Effort)
		fmt.Printf("            RuleSet: \"%s\",\n", issue.RuleSet)
		fmt.Printf("            Rule: \"%s\",\n", issue.Rule)
		fmt.Printf("            Incidents: []api.Incident{\n")
		for _, incident := range issue.Incidents {
			fmt.Printf("                {\n")
			fmt.Printf("                    File: \"%s\",\n", incident.File)
			fmt.Printf("                    Line: %d,\n", incident.Line)
			fmt.Printf("                    Message: \"%s\",\n", incident.Message)
			fmt.Printf("                },\n")
		}
		fmt.Printf("            },\n")
		fmt.Printf("        },\n")
	}
	fmt.Printf("    },\n")
	fmt.Printf("    Dependencies: []api.TechDependency{\n")
	for _, dep := range analysis.Dependencies {
		fmt.Printf("        {\n")
		fmt.Printf("            Name: \"%s\",\n", dep.Name)
		fmt.Printf("            Version: \"%s\",\n", dep.Version)
		fmt.Printf("            Provider: \"%s\",\n", dep.Provider)
		fmt.Printf("        },\n")
	}
	fmt.Printf("    },\n")
	fmt.Printf("}\n")
}

func DumpTags(t *testing.T, tc TC, app api.Application) {
	// Preload tags.
	tags := []api.Tag{}
	for _, tagRef := range app.Tags {
		if tagRef.Source == "Analysis" {
			tag, _ := RichClient.Tag.Get(tagRef.ID)
			tags = append(tags, *tag)
		}
	}
	fmt.Printf("## GOT TAGS FOR \"%s\":", tc.Name)
	fmt.Printf("\n[]api.Tag{\n")
	for _, tag := range tags {
		fmt.Printf("    {Name: \"%s\", Category: api.Ref{Name: \"%s\"}},\n", tag.Name, tag.Category.Name)
	}
	fmt.Printf("}\n")
}
