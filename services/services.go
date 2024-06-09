package services

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/mustafa-a-khan/restapi/models"
)

func Test(Req models.ScrapeRequest) {
	url := Req.URL
	fmt.Println("hii")
	flag := false
	count := 0
	path1 := Req.Parameters["path1"]
	// path2 := Req.Parameters["path2"]

	// Initialize the collector with additional configurations
	c := colly.NewCollector(
		colly.AllowedDomains(url),
		colly.MaxDepth(2), // Limit the depth of links followed
	)

	// Set a request timeout
	c.SetRequestTimeout(30 * time.Second)

	// Limit the number of concurrent requests and set a delay
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		fmt.Println("link:", link)
		h.Request.Visit(link)
	})

	c.OnHTML("title", func(h *colly.HTMLElement) {
		fmt.Println("page title:", h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting:", r.URL)
		if flag {
			r.Abort()
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error:", err)
		// Stop the collector on error
	})
	if count > 10 {
		flag = true
	}
	count = count + 1
	c.Visit("https://" + url + "/" + path1)

}
