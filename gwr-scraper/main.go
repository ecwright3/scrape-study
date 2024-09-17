package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("scrapethissite.com", "www.scrapethissite.com"),
	)

	c.OnHTML(".col-md-4.country", func(e *colly.HTMLElement) {

		e.ChildText(".country-name")
		fmt.Println(strings.TrimSpace(
			e.ChildText(".country-name")),
			strings.TrimSpace(e.ChildText(".country-capital")))

	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	c.Visit("https://www.scrapethissite.com/pages/simple/")
}
