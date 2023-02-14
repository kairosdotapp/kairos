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
			Date:    t0117,
			Account: "time:cust:a:proj1",
			Logs:    []string{"meeting", "work on updated rule"},
			Hours:   2,
		},
		{
			Date:    t0117,
			Account: "time:bec:admin",
			Logs:    []string{},
			Hours:   5,
		},
		{
			Date:    t0117,
			Account: "time:cust:c",
			Logs:    []string{"debug build issues"},
			Hours:   0.75,
		},
		{
			Date:    t0118,
			Account: "time:cust:b",
			Logs:    []string{"weekly design meeting", "work on metrics", "USB performance"},
			Hours:   3,
		},
		{
			Date:    t0118,
			Account: "time:cust:c",
			Logs:    []string{"project setup"},
			Hours:   0.5,
		},
		{
			Date:    t0118,
			Account: "time:bec:siot:go",
			Logs:    []string{},
			Hours:   2,
		},
		{
			Date:    t0118,
			Account: "time:bec:admin",
			Logs:    []string{},
			Hours:   3.25,
		},
		{
			Date:    t0215,
			Account: "time:cust:a:onsite",
			Logs:    []string{"onsite training"},
			Hours:   8,
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
