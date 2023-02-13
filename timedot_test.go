package main

import (
	"strings"
	"testing"
	"time"
)

var testTimedotData = `
2023-01-17 # Tues
time:cust:a:proj1        2
  # meeting
  # work on updated rule
time:bec:admin           5
time:cust:c              0.75
  # debug build issues

2023-01-18 # Wed
time:cust:b              3
  # weekly design meeting
  # work on metrics
  # USB performance
time:cust:c              0.5
  # project setup
time:bec:siot:go         2
time:bec:admin           3.25

2023-02-15 # Fri
time:cust:a:onsite       8
  # onsite training
`

// the following time entries can be used in tests
var t0117, _ = time.Parse(time.DateOnly, "2023-01-17")
var t0118, _ = time.Parse(time.DateOnly, "2023-01-18")
var t0215, _ = time.Parse(time.DateOnly, "2023-02-15")

func TestTimedotParser(t *testing.T) {
	s := strings.NewReader(testTimedotData)

	exp := entries{
		{
			date:    t0117,
			account: "time:cust:a:proj1",
			logs:    []string{"meeting", "work on updated rule"},
			hours:   2,
		},
		{
			date:    t0117,
			account: "time:bec:admin",
			logs:    []string{},
			hours:   5,
		},
		{
			date:    t0117,
			account: "time:cust:c",
			logs:    []string{"debug build issues"},
			hours:   0.75,
		},
		{
			date:    t0118,
			account: "time:cust:b",
			logs:    []string{"weekly design meeting", "work on metrics", "USB performance"},
			hours:   3,
		},
		{
			date:    t0118,
			account: "time:cust:c",
			logs:    []string{"project setup"},
			hours:   0.5,
		},
		{
			date:    t0118,
			account: "time:bec:siot:go",
			logs:    []string{},
			hours:   2,
		},
		{
			date:    t0118,
			account: "time:bec:admin",
			logs:    []string{},
			hours:   3.25,
		},
		{
			date:    t0215,
			account: "time:cust:a:onsite",
			logs:    []string{"onsite training"},
			hours:   8,
		},
	}

	p := newTimedotParser(s)

	entries, err := p.scan()
	if err != nil {
		t.Fatal("Error scanning: ", err)
	}

	if len(entries) != len(exp) {
		t.Fatalf("Did not get the correct # of entries, exp %v, got %v", len(exp), len(entries))
	}

	compareEntries(t, exp, entries)
}
