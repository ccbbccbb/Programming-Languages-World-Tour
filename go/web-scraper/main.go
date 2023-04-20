package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	scrapeUrl := "https://boards.greenhouse.io/paradigm62"

	c := colly.NewCollector(colly.AllowedDomains("greenhouse.io", "www.greenhouse.io", "boards.greenhouse.io"))

	c.OnHTML("div.opening a", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	c.Visit(scrapeUrl)
}