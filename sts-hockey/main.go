package main

import (
	"fmt"
	"strconv"
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
		nxtUrl := e.ChildAttr("li:nth-child(2) > a", "href")
		maxPage, err := strconv.Atoi(e.ChildText("li:nth-last-child(2) > a"))
		if err != nil {
			panic(err)
		}
		pageNum, err := strconv.Atoi(e.ChildText("strong"))

		if err != nil {
			panic(err)
		}
		pageNum++

		lastPage := e.ChildText("li:nth-last-child(2) > a")
		fmt.Println("page number:", strings.TrimSpace(currentPage))
		fmt.Println("last page number:", strings.TrimSpace(lastPage))
		fmt.Println("next page:", pageNum)
		fmt.Println("next url:", nxtUrl)

		if pageNum <= maxPage {
			nxtPage := strconv.Itoa(pageNum)

			link := "https://www.scrapethissite.com/pages/forms/?page_num=" + nxtPage

			e.Request.Visit(link)
			fmt.Println("Visiting2", link)
		}

	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting2", request.URL.String())

	})

	c.Visit("https://www.scrapethissite.com/pages/forms/?page_num=1")
}
