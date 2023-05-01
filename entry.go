package main

import (
	"fmt"
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

func (e entry) String() string {
	return fmt.Sprintf("%v %v %v %v %v", e.Date, e.Account, e.Hours, e.User, e.Cost)
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
	return !e.Date.IsZero()
}

func (e *entry) hasAccount() bool {
	return e.Account != ""
}

func (e *entry) clearDate() {
	e.Date = time.Time{}
	e.Account = ""
	e.Logs = []string{}
	e.Hours = 0
}

type entries []entry

func (es entries) String() string {
	var ret string
	for _, e := range es {
		ret += e.String() + "\n"
	}
	return ret
}

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
			e.Account = strings.TrimPrefix(e.Account, account)
			e.Account = strings.TrimPrefix(e.Account, ":")
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

type projectTotal struct {
	Name  string
	Hours float32
}

type projectTotals []projectTotal

func (es *entries) projectTotals() projectTotals {
	t := make(map[string]float32)
	for _, e := range *es {
		t[e.Account] = t[e.Account] + e.Hours
	}

	ret := make(projectTotals, len(t))

	i := 0
	for k, v := range t {
		if k == "" {
			k = "default"
		}
		if k == "nc" {
			k = "no charge (nc)"
		}
		ret[i] = projectTotal{k, v}
		i++
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

// fuctions for sort
func (es entries) Len() int           { return len(es) }
func (es entries) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es entries) Less(i, j int) bool { return es[i].Date.Before(es[j].Date) }
