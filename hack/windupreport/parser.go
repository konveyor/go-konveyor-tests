package windupreport

import (
	"fmt"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/konveyor/tackle2-hub/api"
)



func Parse(reader io.Reader) (analysis api.Analysis)  {
	  doc, err := goquery.NewDocumentFromReader(reader)
	  if err != nil {
		log.Fatal(err)
	  }
	  log.Printf("%+v", doc)

	  // Story points
	  doc.Find(".effortPoints.total span.points").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		points := s.Text()
		fmt.Printf("Review %d: %s\n", i, points)
	})

	  return
}

func ScrapeReport(reportUrl string) {
	c := colly.NewCollector() 
	c.Visit(reportUrl) 

	// Each application
	c.OnHTML("li.product", func(e *colly.HTMLElement) { 
		
 
		url := e.ChildAttr("a", "href") 
		name := e.ChildText(".name") 
 
		// append 
	})

	// Parse application/dependency

}