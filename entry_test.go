package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/kr/pretty"
)

func testEntries(t *testing.T) entries {
	s := strings.NewReader(testTimedotData)

	p := newTimedotParser(s)

	entries, err := p.scan()
	if err != nil {
		t.Fatal("Error scanning: ", err)
	}

	return entries
}

func testEntriesWithCost(t *testing.T) entries {
	entries := testEntries(t)

	rates, err := parseRates(strings.NewReader(testRateData))
	if err != nil {
		t.Fatal("Error parsing rates: ", err)
	}

	err = entries.populateCost(rates, "cbrake")
	if err != nil {
		t.Fatal("Error populating cost: ", err)
	}

	return entries
}

func compareEntries(t *testing.T, exp, got entries) {
	if len(exp) != len(got) {
		t.Fatalf("not same number of enries, exp: %v, got: %v", len(exp), len(got))
	}
	if !reflect.DeepEqual(exp, got) {
		for i := range exp {
			if !reflect.DeepEqual(exp[i], got[i]) {
				fmt.Println("Failed at index: ", i)
				pretty.Println("exp: ", exp[i], exp[i].Date.Format(time.DateOnly))
				pretty.Println("got: ", got[i], got[i].Date.Format(time.DateOnly))
			}
		}
		t.Fatal("did not get expected result")
	}

}

func TestPopulateCost(t *testing.T) {
	entries := testEntriesWithCost(t)

	if entries[0].User != "cbrake" {
		t.Fatal("User is not correct")
	}

	if entries[0].Cost != 160 {
		t.Fatalf("Cost is not correct, exp %v, got %v", 160,
			entries[0].Cost)
	}
}

func TestFilterAccount(t *testing.T) {
	e := testEntries(t)

	filtered := e.filterAccount("time:cust:a")

	exp := entries{
		{
			Date:    t0117,
			Account: "time:cust:a:proj1",
			Logs:    []string{"meeting", "work on updated rule"},
			Hours:   2,
		},
		{
			Date:    t0215,
			Account: "time:cust:a:onsite",
			Logs:    []string{"onsite training"},
			Hours:   8,
		},
	}
	compareEntries(t, exp, filtered)
}

func TestFilterDate(t *testing.T) {
	e := testEntries(t)

	start, _ := time.Parse(time.DateOnly, "2023-02-15")
	end, _ := time.Parse(time.DateOnly, "2023-02-15")

	filtered := e.filterDate(start, end)

	exp := entries{
		{
			Date:    t0215,
			Account: "time:cust:a:onsite",
			Logs:    []string{"onsite training"},
			Hours:   8,
		},
	}

	compareEntries(t, exp, filtered)
}

func TestInvoice(t *testing.T) {
	entries := testEntriesWithCost(t)

	invoice, err := entries.invoice()
	if err != nil {
		t.Fatal("Error creating invoice: ", err)
	}

	os.WriteFile("invoice.html", []byte(invoice), 0644)
}
