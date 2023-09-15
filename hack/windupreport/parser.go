package windupreport

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/binding"
	"github.com/konveyor/tackle2-hub/test/api/client"
	"github.com/konveyor/tackle2-hub/test/assert"
)

// Setup Hub API client
var RichClient *binding.RichClient

func init() {
	// Prepare RichClient and login to Hub API (configured from env variables).
	RichClient = client.PrepareRichClient()
}

func matchIncident(re *regexp.Regexp, rawIncident string) (match string) {
	matches := re.FindStringSubmatch(rawIncident)
	if len(matches) > 1 {
		match = matches[len(matches)-1]
	} else {
		fmt.Printf("Cannot match Incident %s with %v\n", rawIncident, re)
	}
	return
}

func findRuleset(t *testing.T, appId uint, ruleName string) (ruleset string) {
	doc, err := goquery.NewDocumentFromReader(downloadReport(t, appId, "/windup/report/reports/windup_ruleproviders.html"))
	assert.Must(t, err)
	rulesetSection := doc.Find("a.anchor[name="+ruleName+"]").Closest("div.panel-primary")
	ruleset = rulesetSection.Find("h3.panel-title").Text()
	return
}

func populateIncidents(t *testing.T, appId uint, appDetailPath string, issues *[]api.Issue) (incidentedIssues []api.Issue) {
	doc, err := goquery.NewDocumentFromReader(downloadReport(t, appId, "/windup/report/reports/"+appDetailPath))
	assert.Must(t, err)

	incidentedIssues = *issues

	// radek po radku podle jemna issue rozklikavat soubory a pridavat jednotlive issue
	doc.Find("tr.projectFile > td.path > a").Each(func(i int, s *goquery.Selection) {
		//fileName := strings.TrimSpace(s.Text())
		filePath := strings.Split(s.AttrOr("href", ""), "?")
		fileName := strings.ReplaceAll(filePath[0], ".html", "")
		fileName = strings.ReplaceAll(fileName, "_", ".")
		detailDoc, err := goquery.NewDocumentFromReader(downloadReport(t, appId, "/windup/report/reports/"+filePath[0]))
		assert.Must(t, err)

		// Source code LINE incidents - parsing javascript/jQuery code line-by-line
		for _, incidentLine := range strings.Split(detailDoc.Text(), "\n") {
			if !strings.HasPrefix(incidentLine, `$("<a name`) {
				continue
			}

			reDescription := regexp.MustCompile(`<div class='inline-comment-body'>(.*)</div>"`)
			reIssueName := regexp.MustCompile(`<div class='inline-comment-heading'><strong class='[a-z ]+'>(.*)</strong>`)
			reLine := regexp.MustCompile(`'\#([0-9]+)\-inlines'`)
			reRule := regexp.MustCompile(`'windup_ruleproviders.html#([A-Za-z0-9-]+)'`)

			issueName := stripTags(matchIncident(reIssueName, incidentLine))
			ruleName := matchIncident(reRule, incidentLine)
			rulesetName := findRuleset(t, appId, ruleName)

			incident := api.Incident{}
			incident.File = fileName
			incident.Line, _ = strconv.Atoi(matchIncident(reLine, incidentLine))
			incident.Message = stripTags(matchIncident(reDescription, incidentLine))
			//fmt.Printf("FOUND LINE INCIDENT: %+v\n", incident)

			// Append Incident to the parent Issue.
			for i, issue := range incidentedIssues {
				if issue.Name == issueName {
					incidentedIssues[i].Incidents = append(incidentedIssues[i].Incidents, incident)
					if issue.Rule == "" {
						incidentedIssues[i].Rule = ruleName
						incidentedIssues[i].RuleSet = rulesetName
					}
				}
			}
		}

		// Whole FILE incidents, e.g. information
		detailDoc.Find("ul.classifications > li").Each(func(i int, s *goquery.Selection) {
			issueName := s.Find("div.title em").Text()
			ruleName := strings.ReplaceAll(s.Find("div.title a").AttrOr("title", ""), "View Rule: ", "")
			rulesetName := findRuleset(t, appId, ruleName)

			incident := api.Incident{}
			incident.File = fileName
			incident.Message = s.Find("div.desc").Text()
			//fmt.Printf("FOUND FILE INCIDENT: %+v\n", incident)

			// Append Incident to the parent Issue.
			for i, issue := range incidentedIssues {
				if issue.Name == issueName {
					incidentedIssues[i].Incidents = append(incidentedIssues[i].Incidents, incident)
					if issue.Rule == "" {
						incidentedIssues[i].Rule = ruleName
						incidentedIssues[i].RuleSet = rulesetName
					}
				}
			}
		})

	})
	return
}

func ParseWindupReportIssues(t *testing.T, appId uint) (issues []api.Issue) {
	doc, err := goquery.NewDocumentFromReader(downloadReport(t, appId, "/windup/report/reports/migration_issues.html"))
	assert.Must(t, err)

	// Populate categories
	categories := []string{}
	doc.Find("table.tablesorter.migration-issues-table").Each(func(i int, s *goquery.Selection) {
		catName := strings.ReplaceAll(s.AttrOr("id", ""), "table-", "")
		categories = append(categories, catName)
	})

	// Find issues by category.
	for _, category := range categories {	// Common are: cloud-mandatory, information
		doc.Find(fmt.Sprintf("#table-%s tr.problemSummary", category)).Each(func(i int, s *goquery.Selection) {
			issue := api.Issue{}
			issue.Category = category
			issue.Description = s.Find("td:nth-child(5n+4)").Text()
			issue.Effort, _ = strconv.Atoi(s.Find("td:nth-child(5n+5)").Text())
			issue.Name = strings.TrimSpace(s.Find("td:first-child").Text())
			issues = append(issues, issue)
		})
	}

	// Populate incidents from application detail page.
	appDetailPath := strings.Replace(doc.Find("ul.nav > li:nth-child(2) > a").AttrOr("href", ""), "report_index_", "ApplicationDetails_", 1)
	issues = populateIncidents(t, appId, appDetailPath, &issues)

	// Sort issues to have stable order.
	sort.SliceStable(issues, func(i, j int) bool {
		return issues[i].Category+issues[i].Name < issues[j].Category+issues[j].Name
	})
	return
}

func Parse(t *testing.T, appId uint) (analysis api.Analysis) {
	doc, err := goquery.NewDocumentFromReader(downloadReport(t, appId, "/windup/report/index.html"))
	assert.Must(t, err)

	doc.Find("div.stats").Each(func(i int, s *goquery.Selection) {
		analysis.Effort, _ = strconv.Atoi(s.Find("span.points").Text())
	})

	analysis.Issues = ParseWindupReportIssues(t, appId)

	return
}

// Get report file and return reader.
func downloadReport(t *testing.T, appId uint, path string) (reader *bufio.Reader) {
	// Prepare temp directory.
	tempDir, err := os.MkdirTemp("/tmp", fmt.Sprintf("app_%d_windup_report", appId))
	if err != nil {
		panic(err.Error())
	}
	fileName := filepath.Join(tempDir, filepath.Base(path))
	// Download the file from bucket.
	err = RichClient.Application.Bucket(appId).Get(path, tempDir)
	assert.Must(t, err)
	f, err := os.Open(fileName)
	assert.Must(t, err)
	// Return reader.
	reader = bufio.NewReader(f)
	return
}

func stripTags(input string) (output string) {
	output = regexp.MustCompile(`<.*?>`).ReplaceAllString(input, "")
	return
}