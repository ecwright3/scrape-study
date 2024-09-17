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

	c.OnHTML(".team", func(e *colly.HTMLElement) {
		if strings.Contains(strings.ToLower(e.ChildText(".name")), "capitals") {
			fmt.Println(strings.TrimSpace(
				e.ChildText(".name")),
				strings.TrimSpace(e.ChildText(".year")),
				strings.TrimSpace(e.ChildText(".pct")))
		}

	})

	c.OnHTML(".pagination", func(e *colly.HTMLElement) {
		currentPage := e.ChildText("strong")
		lastPage := e.ChildText("li:nth-last-child(2) > a")
		fmt.Println("page number:", strings.TrimSpace(currentPage))
		fmt.Println("last page number:", strings.TrimSpace(lastPage))
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())

	})
	c.Visit("https://www.scrapethissite.com/pages/forms/?page_num=20")
}
