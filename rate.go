package main

import (
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
)

type rate struct {
	Account string  `csv:"account"`
	User    string  `csv:"user"`
	Rate    float32 `csv:"rate"`
}

func parseRates(r io.Reader) ([]rate, error) {
	var rates []rate

	if err := gocsv.Unmarshal(r, &rates); err != nil {
		return rates, fmt.Errorf("csv error: %v", err)
	}

	return rates, nil
}
