package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func ScrapeLinks(link string) []string {
	var hrefs []string
	c := colly.NewCollector()
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		hrefs = append(hrefs, e.ChildAttrs("a", "href")...)
	})

	c.Visit(link)
	return hrefs
}

func CheckErr(err error) bool {
	if err != nil {
		log.Fatalf("Couldn't create file, err: %q", err)
		return true
	}
	return false
}

func main() {
	link := "https://en.wikipedia.org/wiki/Web_scraping"

	fName := "Links.txt"
	file, err := os.Create(fName)

	if CheckErr(err) {
		return
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(ScrapeLinks(link))

	fName = "table.csv"
	file, err = os.Create(fName)

	c := colly.NewCollector()

	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Scrapping Complete")
	})

	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}
