# Cryptocurrency Scraper

This is a Go program that scrapes cryptocurrency data from Yahoo Finance.

### Installation

- Make sure you have Go installed and configured on your system. You can check the installation by running `go version` in your terminal.
- Clone this repository using `git clone https://github.com/dipankarupd/crypto-scraper.git`.
- Run the program using `go run main.go`.

### Dependencies

This program uses the following Go libraries:

- `github.com/gocolly/colly/` - This is a popular web scraping library for Go.

### How it Works

- The program first retrieves a list of crypto symbols from Yahoo Finance using the `util.GetScrapedSymbols `function.

- It then iterates over the symbols and visits the corresponding Yahoo Finance page for each symbol using the `colly` library.

- While scraping the page, the program extracts the company name, price, and change in price using CSS selectors. This data is stored in a `Crypto` struct.

- Finally, all the scraped data is written to a CSV file named `result.csv` using the `util.AddToCSV` function.

## Running the program

Compile and run the program using the following command:

```
go run main.go
```

#### Blog

You can read about the comprehensive description on how the project is made via this blog.<br> [Blog](https://medium.com/@drupd17/web-scraping-with-golang-scrape-cryptocurrency-data-with-go-e470eb9e916e)
