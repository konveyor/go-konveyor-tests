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

var (
	// Setup Hub API client
	Client     *binding.Client
	RichClient *binding.RichClient

	// Analysis waiting loop 5 minutes (60 * 5s)
	Retry = 100
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
	// Analysis parameters.
	Task     api.Task
	TaskData string
	Sources  []string
	Targets  []string
	Labels   addon.Labels
	Rules    addon.Rules
	WithDeps bool
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
	fmt.Printf("}\n")
}

func DumpTags(t *testing.T, tc TC, app api.Application) {
	fmt.Printf("## GOT TAGS FOR \"%s\":", tc.Name)
	fmt.Printf("\n[]api.Tag{\n")
	for _, tag := range app.Tags {
		if tag.Source == "Analysis" {
			fmt.Printf("    {Name: \"%s\"},\n", tag.Name)
		}
	}
	fmt.Printf("}\n")
}
