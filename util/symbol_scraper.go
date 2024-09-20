package util

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func GetScrapedSymbols() []string {
	c := colly.NewCollector()

	var symbols []string

	c.OnHTML("tr td span.symbol", func(h *colly.HTMLElement) {
		symbol := h.Text
		symbols = append(symbols, symbol)
	})

	if err := c.Visit("https://finance.yahoo.com/crypto"); err != nil {
		fmt.Println("Error visiting the page: ", err)
	}

	return symbols
}
