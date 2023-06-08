package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

type invoiceData struct {
	Date          string
	Number        int
	Entries       invoiceEntries
	Hours         float32
	Cost          float32
	Customer      customer
	ProjectTotals projectTotals
}

type invoiceEntry struct {
	Date    string
	Project string
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
		ret[i].Project = e.Account
		ret[i].Logs = e.Logs
		ret[i].Hours = e.Hours
		ret[i].User = e.User
		ret[i].Cost = e.Cost
	}

	return ret
}

// invoiceMonth is given in form of YYYY-MM
// if date is "", then current date is used
func invoice(number int, es entries, custs customers, account string, start, end time.Time,
	date, tplFile string) (string, error) {
	accountEntries := es.filterAccount(account)
	accountEntries = accountEntries.filterDate(start, end)

	projectTotals := accountEntries.projectTotals()

	cust, ok := custs.find(account)
	if !ok {
		return "", fmt.Errorf("Did not find customer")
	}

	f, err := os.ReadFile(tplFile)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v", err)
	}

	t := template.New("invoice")

	t, err = t.Parse(string(f))
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v", err)
	}

	var ret strings.Builder

	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}

	data := invoiceData{
		Number:        number,
		Date:          date,
		Customer:      cust,
		Entries:       newInvoiceEntries(accountEntries),
		Cost:          accountEntries.cost(),
		Hours:         accountEntries.hours(),
		ProjectTotals: projectTotals,
	}

	err = t.Execute(&ret, data)
	if err != nil {
		return "", fmt.Errorf("Error applying template: %v", err)
	}

	return ret.String(), nil
}
