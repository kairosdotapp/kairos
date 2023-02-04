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
}

func newParser(scanner *bufio.Scanner) *parser {
	return &parser{scanner: scanner}
}

func (p *parser) scan() (*entry, error) {
	for p.scanner.Scan() {
		t := p.scanner.Text()
		c := reDate.FindString(t)
		if len(c) > 0 {fmt.Println("Date matches: ", c)}
	}

	if err := p.scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}
