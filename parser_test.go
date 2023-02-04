package main

import (
	"bufio"
	"strings"
	"testing"
)

var testData = `
2023-01-17 # Tues
time:cust:a              2
  # weekly priority meeting
  # work on updated rule offline design
  # portal: update UI to create offline comp IO conditions
time:bec:admin           5
time:cust:c              0.75
  # work on debugging OE build issues

2023-01-18 # Wed
time:cust:b              3
  # weekly design meeting
  # work on system/app metrics
  # work on USB performance
time:cust:c              0.5
  # project setup, work on OE build issues
time:bec:siot:go         2
time:bec:admin           3.25
`

func TestParser(t *testing.T) {
	s := strings.NewReader(testData)

	scanner := bufio.NewScanner(s)

	p := newParser(scanner)

	var entries []entry

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

	if len(entries) != 2 {
		t.Fatal("Expected 2 entries, got: ", len(entries))
	}

}
