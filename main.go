package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	articles := []string{}

	// Find articles
	c.OnHTML("article > h2 > a", func(e *colly.HTMLElement) {
		articles = append(articles, e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://datoga.es")

	fmt.Printf("We have found %d articles\n", len(articles))

	for i, article := range articles {
		fmt.Printf("Article %d: %s\n", (i + 1), article)
	}
}
