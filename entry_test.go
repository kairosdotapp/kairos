package main

import (
	"strings"
	"testing"
)

func TestPopulateCost(t *testing.T) {
	s := strings.NewReader(testTimedotData)

	p := newTimedotParser(s)

	entries, err := p.scan()
	if err != nil {
		t.Fatal("Error scanning: ", err)
	}

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
