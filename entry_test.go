package main

import (
	"fmt"
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

func compareEntries(t *testing.T, exp, got entries) {
	if len(exp) != len(got) {
		t.Fatalf("not same number of enries, exp: %v, got: %v", len(exp), len(got))
	}
	if !reflect.DeepEqual(exp, got) {
		for i := range exp {
			if !reflect.DeepEqual(exp[i], got[i]) {
				fmt.Println("Failed at index: ", i)
				pretty.Println("exp: ", exp[i], exp[i].date.Format(time.DateOnly))
				pretty.Println("got: ", got[i], got[i].date.Format(time.DateOnly))
			}
		}
		t.Fatal("did not get expected result")
	}

}

func TestPopulateCost(t *testing.T) {
	entries := testEntries(t)

	rates, err := parseRates(strings.NewReader(testRateData))
	if err != nil {
		t.Fatal("Error parsing rates: ", err)
	}

	err = entries.populateCost(rates, "cbrake")
	if err != nil {
		t.Fatal("Error populating cost: ", err)
	}

	if entries[0].user != "cbrake" {
		t.Fatal("User is not correct")
	}

	if entries[0].cost != 160 {
		t.Fatalf("Cost is not correct, exp %v, got %v", 160,
			entries[0].cost)
	}
}

func TestFilterAccount(t *testing.T) {
	e := testEntries(t)

	filtered := e.filterAccount("time:cust:a")

	exp := entries{
		{
			date:    t0117,
			account: "time:cust:a:proj1",
			logs:    []string{"meeting", "work on updated rule"},
			hours:   2,
		},
		{
			date:    t0215,
			account: "time:cust:a:onsite",
			logs:    []string{"onsite training"},
			hours:   8,
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
			date:    t0215,
			account: "time:cust:a:onsite",
			logs:    []string{"onsite training"},
			hours:   8,
		},
	}

	compareEntries(t, exp, filtered)
}
