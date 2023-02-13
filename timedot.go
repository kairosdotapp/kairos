package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type timedotEntry struct {
	date    string
	account string
	logs    []string
	hours   float32
}

func (e *timedotEntry) setDate(d string) {
	e.date = d
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

func (e *timedotEntry) setAccount(a string) {
	e.account = a
	e.logs = []string{}
	e.hours = 0
}

func (e *timedotEntry) hasDate() bool {
	if e.date != "" {
		return true
	}

	return false
}

func (e *timedotEntry) hasAccount() bool {
	if e.account != "" {
		return true
	}

	return false
}

func (e *timedotEntry) clearDate() {
	e.date = ""
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

func (e *timedotEntry) clearAccount() {
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

type timedotEntrys []timedotEntry

var reDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)
var reAccount = regexp.MustCompile(`^(time\S*)\s*(([0-9]*[.])?[0-9]+)`)
var reLog = regexp.MustCompile(`^  #\s*(\S+.*)`)

type timedotParserState int

const (
	parserStateNone timedotParserState = iota
	parserStateEntry
)

type timedotParser struct {
	scanner      *bufio.Scanner
	state        timedotParserState
	currentEntry timedotEntry
}

func newTimedotParser(r io.Reader) *timedotParser {
	return &timedotParser{scanner: bufio.NewScanner(r)}
}

func (p *timedotParser) scan() (timedotEntrys, error) {
	var ret timedotEntrys

	for {
		e, err := p.scanEntry()

		if err != nil {
			return nil, err
		}

		if e == nil {
			break
		}

		ret = append(ret, *e)
	}

	return ret, nil
}

func (p *timedotParser) scanEntry() (*timedotEntry, error) {
	for p.scanner.Scan() {
		t := p.scanner.Text()

		d := reDate.FindString(t)
		if len(d) > 0 {
			// we found a new entry, do we have a current one?
			ret := p.currentEntry
			p.currentEntry.setDate(d)

			if ret.hasDate() {
				return &ret, nil
			}
			continue
		}

		aMatches := reAccount.FindStringSubmatch(t)
		if len(aMatches) >= 2 {
			a := aMatches[1]
			ret := p.currentEntry
			p.currentEntry.setAccount(a)

			if len(aMatches) >= 3 {
				hS := aMatches[2]
				h, err := strconv.ParseFloat(hS, 32)
				if err != nil {
					return nil, fmt.Errorf("Error parsing hours: %v", err)
				}
				p.currentEntry.hours = float32(h)
			}

			if ret.hasAccount() {
				return &ret, nil
			}

			continue
		}

		lMatches := reLog.FindStringSubmatch(t)
		if len(lMatches) >= 2 {
			p.currentEntry.logs = append(p.currentEntry.logs, lMatches[1])
		}
	}

	if p.currentEntry.hasDate() {
		r := p.currentEntry
		p.currentEntry.clearDate()
		return &r, nil
	}

	if err := p.scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}
