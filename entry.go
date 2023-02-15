package main

import (
	"strings"
	"time"
)

type entry struct {
	Date    time.Time
	Account string
	Logs    []string
	Hours   float32
	User    string
	Cost    float32
}

func (e *entry) setDate(d time.Time) {
	e.Date = d
	e.Account = ""
	e.Logs = []string{}
	e.Hours = 0
}

func (e *entry) setAccount(a string) {
	e.Account = a
	e.Logs = []string{}
	e.Hours = 0
}

func (e *entry) hasDate() bool {
	if !e.Date.IsZero() {
		return true
	}

	return false
}

func (e *entry) hasAccount() bool {
	if e.Account != "" {
		return true
	}

	return false
}

func (e *entry) clearDate() {
	e.Date = time.Time{}
	e.Account = ""
	e.Logs = []string{}
	e.Hours = 0
}

func (e *entry) clearAccount() {
	e.Account = ""
	e.Logs = []string{}
	e.Hours = 0
}

type entries []entry

func (es *entries) populateCost(r rates, user string) error {
	for i, e := range *es {
		(*es)[i].Cost = r.find(e.Account, user) * e.Hours
		(*es)[i].User = user
	}

	return nil
}

// returns entries where account param matches beginning of entry account field
func (es *entries) filterAccount(account string) entries {
	var ret entries

	for _, e := range *es {
		if strings.HasPrefix(e.Account, account) {
			ret = append(ret, e)
		}
	}

	return ret
}

func (es *entries) filterDate(start, end time.Time) entries {
	var ret entries

	for _, e := range *es {
		if e.Date.Before(start) {
			continue
		}

		if e.Date.After(end) {
			continue
		}

		ret = append(ret, e)
	}
	return ret
}

func (es *entries) cost() float32 {
	var ret float32
	for _, e := range *es {
		ret += e.Cost
	}
	return ret
}

func (es *entries) hours() float32 {
	var ret float32
	for _, e := range *es {
		ret += e.Hours
	}
	return ret
}
