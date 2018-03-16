package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "./export/tokenmarket.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Status", "Name", "Symbol", "Description"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("table.table-listing tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".asset-status"),
			e.ChildText(".col-asset-name"),
			e.ChildText(".col-asset-symbol"),
			e.ChildText(".col-asset-description"),
		})
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://tokenmarket.net/blockchain/all-assets?batch_num=0&batch_size=1000")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
