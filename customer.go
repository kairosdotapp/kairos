package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/gocarina/gocsv"
)

type customer struct {
	Account string `csv:"account"`
	Name    string `csv:"name"`
	Address string `csv:"address"`
	City    string `csv:"city"`
	State   string `csv:"state"`
	Zip     string `csv:"zip"`
}

func (c *customer) match(account string) bool {
	if strings.HasPrefix(account, c.Account) {
		return true
	}

	return false
}

type customers []customer

func (cs *customers) find(account string) (customer, bool) {
	for _, c := range *cs {
		if c.match(account) {
			return c, true
		}
	}

	return customer{}, false
}

func parseCustomers(r io.Reader) (customers, error) {
	var customers customers

	if err := gocsv.Unmarshal(r, &customers); err != nil {
		return customers, fmt.Errorf("csv error: %v", err)
	}

	return customers, nil
}
