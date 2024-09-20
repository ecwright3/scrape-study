package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	//Create New Collector
	c := colly.NewCollector(
		colly.AllowedDomains("nfl.com", "www.nfl.com", "gsm-widgets.betstream.betgenius.com"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:130.0) Gecko/20100101 Firefox/130.0"),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1, //Adjust number of concurrent requests
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(".css-jb7qf-Column", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText(".css-text-146c3p1"))
		fmt.Println(e.Text)
	})

	c.Visit("https://www.nfl.com/scores/")

}

//Request URL: https://gsm-widgets.betstream.betgenius.com/widget-data/mainlineswid
//content-type: application/json; charset=utf-8
//class="css-view-175oi2r r-borderRightColor-dhdqoo r-borderRightWidth-13l2t4g r-padding-nsbfu8 r-width-6gcxwl"
