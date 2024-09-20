package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	start := time.Now()
	//Create Collector
	c := colly.NewCollector(
		colly.AllowedDomains("scrapethissite.com", "www.scrapethissite.com"),
		colly.Async(true), //Enable asynchronous requests
	)

	// Limit number of concrrent requests
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 6, //Adjust number of concurrent requests
	})

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

		//check pages concurrently
		for pageNum <= maxPage {
			nxtPage := strconv.Itoa(pageNum)

			link := "https://www.scrapethissite.com/pages/forms/?page_num=" + nxtPage
			e.Request.Visit(link)
			fmt.Println("Visiting2", link)
			pageNum++
		}

	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())

	})

	c.Visit("https://www.scrapethissite.com/pages/forms/?page_num=1")

	// Wait until threads are finished
	c.Wait()
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

//https://my.peachjar.com/explore/all?audienceId=all&tab=school&districtId=50183&audienceType=school
//https://my.peachjar.com/explore/all?audienceId=all&tab=community&districtId=50183&audienceType=school
