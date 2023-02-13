package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
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

func TestTimedotParser(t *testing.T) {
	s := strings.NewReader(testTimedotData)

	exp := []timedotEntry{
		{
			date:    "2023-01-17",
			account: "time:cust:a:proj1",
			logs:    []string{"meeting", "work on updated rule"},
			hours:   2,
		},
		{
			date:    "2023-01-17",
			account: "time:bec:admin",
			logs:    []string{},
			hours:   5,
		},
		{
			date:    "2023-01-17",
			account: "time:cust:c",
			logs:    []string{"debug build issues"},
			hours:   0.75,
		},
		{
			date:    "2023-01-18",
			account: "time:cust:b",
			logs:    []string{"weekly design meeting", "work on metrics", "USB performance"},
			hours:   3,
		},
		{
			date:    "2023-01-18",
			account: "time:cust:c",
			logs:    []string{"project setup"},
			hours:   0.5,
		},
		{
			date:    "2023-01-18",
			account: "time:bec:siot:go",
			logs:    []string{},
			hours:   2,
		},
		{
			date:    "2023-01-18",
			account: "time:bec:admin",
			logs:    []string{},
			hours:   3.25,
		},
		{
			date:    "2023-02-15",
			account: "time:cust:a:onsite",
			logs:    []string{"onsite training"},
			hours:   8,
		},
	}

	scanner := bufio.NewScanner(s)

	p := newTimedotParser(scanner)

	var entries []timedotEntry

	for {
		e, err := p.scan()

		if err != nil {
			t.Fatal(err)
		}

		if e == nil {
			break
		}

		entries = append(entries, *e)
	}

	if len(entries) != len(exp) {
		t.Fatalf("Did not get the correct # of entries, exp %v, got %v", len(exp), len(entries))
	}

	if !reflect.DeepEqual(exp, entries) {
		for i := range exp {
			if !reflect.DeepEqual(exp[i], entries[i]) {
				fmt.Println("Failed at index: ", i)
				pretty.Println("exp: ", exp[i])
				pretty.Println("entries: ", entries[i])
			}
		}
		t.Fatal("did not get expected result")
	}

}
