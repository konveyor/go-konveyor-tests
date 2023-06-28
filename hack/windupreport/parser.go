package windupreport

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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

func ParseWindupReportIssues(t *testing.T, appId uint) (issues []api.Issue) {
	doc, err := goquery.NewDocumentFromReader(downloadReport(t, appId, "/windup/report/reports/migration_issues.html"))
	assert.Must(t, err)

	for _, category := range []string{"cloud-mandatory", "information"} {
		doc.Find(fmt.Sprintf("#table-%s tr.problemSummary", category)).Each(func(i int, s *goquery.Selection) {
			issue := api.Issue{}
			issue.Category = category
			issue.Description = s.Find("td:nth-child(5n+4)").Text()
			issue.Effort, _ = strconv.Atoi(s.Find("td:nth-child(5n+5)").Text())
			issue.Name = strings.TrimSpace(s.Find("td:first-child").Text())
			issues = append(issues, issue)
		})
	}
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
