package main

import (
	"bufio"
	"fmt"
	"regexp"
)

type entry struct {
	date string
	who  string
	logs []string
	time float32
}

var reDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)

type parserState int

const (
	parserStateNone parserState = iota
	parserStateEntry
)

type parser struct {
	scanner *bufio.Scanner
	state   parserState
	currentEntry *entry
}

func newParser(scanner *bufio.Scanner) *parser {
	return &parser{scanner: scanner}
}

func (p *parser) scan() (*entry, error) {
	for p.scanner.Scan() {
		t := p.scanner.Text()
		d := reDate.FindString(t)

		if len(d) > 0 {
			// we found a new entry, do we have a current one?
			if p.currentEntry != nil {
				r := *p.currentEntry
				p.currentEntry = &entry{date:d} 
				fmt.Println("returning: ", r)
				return &r, nil
			}

			p.currentEntry = &entry{date:d} 
		}
	}

	if p.currentEntry != nil {
		r := *p.currentEntry
		p.currentEntry = nil
		return &r, nil
	}

	if err := p.scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}
