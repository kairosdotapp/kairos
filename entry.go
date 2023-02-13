package main

import (
	"strings"
	"time"
)

type entry struct {
	date    time.Time
	account string
	logs    []string
	hours   float32
	user    string
	cost    float32
}

func (e *entry) setDate(d time.Time) {
	e.date = d
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

func (e *entry) setAccount(a string) {
	e.account = a
	e.logs = []string{}
	e.hours = 0
}

func (e *entry) hasDate() bool {
	if !e.date.IsZero() {
		return true
	}

	return false
}

func (e *entry) hasAccount() bool {
	if e.account != "" {
		return true
	}

	return false
}

func (e *entry) clearDate() {
	e.date = time.Time{}
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

func (e *entry) clearAccount() {
	e.account = ""
	e.logs = []string{}
	e.hours = 0
}

type entries []entry

func (es *entries) populateCost(r rates, user string) error {
	for i, e := range *es {
		(*es)[i].cost = r.find(e.account, user) * e.hours
		(*es)[i].user = user
	}

	return nil
}

// returns entries where account param matches beginning of entry account field
func (es *entries) filterAccount(account string) entries {
	var ret entries

	for _, e := range *es {
		if strings.HasPrefix(e.account, account) {
			ret = append(ret, e)
		}
	}

	return ret
}

func (es *entries) filterDate(start, end time.Time) entries {
	var ret entries

	for _, e := range *es {
		if e.date.Before(start) {
			continue
		}

		if e.date.After(end) {
			continue
		}

		ret = append(ret, e)
	}
	return ret
}
