package main

import (
	"fmt"

	"github.com/dipankarupd/cryptoscraper/model"
	"github.com/dipankarupd/cryptoscraper/util"
	"github.com/gocolly/colly/v2"
)

func main() {

	symbols := util.GetScrapedSymbols()

	for _, s := range symbols {
		fmt.Println(s)
	}

	cryptos := []model.Crypto{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping ", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error while scraping ", err)
	})

	c.OnHTML("section.container.yf-1s1umie", func(h *colly.HTMLElement) {

		crypto := model.Crypto{}

		crypto.Name = h.ChildText("h1")
		fmt.Println("Company Name: ", crypto.Name)

		crypto.Price = h.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Crypto Price: ", crypto.Price)

		crypto.Change = h.ChildText("fin-streamer[data-field=regularMarketChangePercent]")
		fmt.Println("Value Change: ", crypto.Change)

		cryptos = append(cryptos, crypto)
	})

	c.Wait()

	// visit the crypto sites:
	for _, s := range symbols {
		c.Visit("https://finance.yahoo.com/quote/" + s + "/")
	}

	fmt.Println(cryptos)

	// add the scraped data into CSV file
	util.AddToCSV(cryptos)

}
