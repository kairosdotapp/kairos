package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

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

func TestFilterEntries(t *testing.T) {
	e := testEntries(t)

	filtered := e.filter("time:cust:a")

	exp := entries{
		{
			date:    "2023-01-17",
			account: "time:cust:a:proj1",
			logs:    []string{"meeting", "work on updated rule"},
			hours:   2,
		},
		{
			date:    "2023-02-15",
			account: "time:cust:a:onsite",
			logs:    []string{"onsite training"},
			hours:   8,
		},
	}

	if !reflect.DeepEqual(exp, filtered) {
		for i := range exp {
			if !reflect.DeepEqual(exp[i], e[i]) {
				fmt.Println("Failed at index: ", i)
				pretty.Println("exp: ", exp[i])
				pretty.Println("entries: ", e[i])
			}
		}
		t.Fatal("did not get expected result")
	}

}
