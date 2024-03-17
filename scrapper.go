package main

import (
	"github.com/gocolly/colly"
)

type musicData struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Rank   int    `json:"rank"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.billboard.com"),
	)

	var music []musicData
	rank := 1

	c.OnHTML("div[class=o-chart-results-list-row-container] li.lrv-u-flex-grow-1", func(e *colly.HTMLElement) {
		title := e.ChildText("h3[id=title-of-a-story]")
		artist := e.ChildText("span.a-font-primary-s")

		if title != "" && artist != "" {
			music = append(music, musicData{
				Title:  title,
				Artist: artist,
				Rank:   rank,
			})
			rank++
		}
	})

	// Start scraping
	c.Visit("https://www.billboard.com/charts/hot-100/1991-07-13/")

	// Print data
	for _, m := range music {
		println(m.Rank, m.Title, m.Artist)
	}
}
