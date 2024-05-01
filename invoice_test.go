package main

import (
	"os"
	"testing"
	"time"
)

func TestInvoice(t *testing.T) {
	acnt := "time:cust:a"
	entries := testEntriesWithCost(t)
	custs := testCustomers(t)

	start, _ := time.Parse(time.DateOnly, "2023-01-01")
	end := start.AddDate(0, 1, 0).Add(-time.Second)

	invoice, err := invoice(1055, entries, custs, acnt, start, end, "", "invoice.tpl")
	if err != nil {
		t.Fatal("Error creating invoice: ", err)
	}

	_ = os.WriteFile("invoice.html", []byte(invoice), 0644)
}
