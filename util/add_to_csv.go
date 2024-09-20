package util

import (
	"encoding/csv"
	"os"

	"github.com/dipankarupd/cryptoscraper/model"
)

func AddToCSV(crypto []model.Crypto) {

	file, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{"company", "price", "change"}

	writer.Write(headers)

	// add the crypto data into csv file:

	for _, c := range crypto {
		record := []string{c.Name, c.Price, c.Change}
		writer.Write(record)
	}

	defer writer.Flush()

}
