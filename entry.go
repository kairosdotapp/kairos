package main

import (
	"fmt"
	"html/template"
	"os"
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

type invoiceData struct {
	Number  int
	Entries invoiceEntries
	Hours   float32
	Cost    float32
}

type invoiceEntry struct {
	Date    string
	Account string
	Logs    []string
	Hours   float32
	User    string
	Cost    float32
}

type invoiceEntries []invoiceEntry

func newInvoiceEntries(in entries) invoiceEntries {
	ret := make(invoiceEntries, len(in))

	for i, e := range in {
		ret[i].Date = e.Date.Format(time.DateOnly)
		ret[i].Account = e.Account
		ret[i].Logs = e.Logs
		ret[i].Hours = e.Hours
		ret[i].User = e.User
		ret[i].Cost = e.Cost
	}

	return ret
}

func (es *entries) invoice() (string, error) {
	f, err := os.ReadFile("invoice.tpl")
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v", err)
	}

	t := template.New("invoice")

	t, err = t.Parse(string(f))
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v", err)
	}

	var ret strings.Builder

	data := invoiceData{
		Number:  1023,
		Entries: newInvoiceEntries(*es),
		Cost:    es.cost(),
		Hours:   es.hours(),
	}

	err = t.Execute(&ret, data)
	if err != nil {
		return "", fmt.Errorf("Error applying template: %v", err)
	}

	return ret.String(), nil
}
