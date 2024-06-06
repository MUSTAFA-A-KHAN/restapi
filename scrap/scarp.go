package scrap

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func scrapperService(url string) {

	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		fmt.Print("link:", link)
		h.Request.Visit(link)
	})

	c.OnHTML("title", func(h *colly.HTMLElement) {
		fmt.Println("page title:", h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		println("error")
	})

	c.Visit("url")

}
