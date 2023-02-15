package main

import (
	"os"
	"testing"
)

func TestInvoice(t *testing.T) {
	acnt := "time:cust:a"
	entries := testEntriesWithCost(t)
	custs := testCustomers(t)

	invoice, err := invoice(entries, custs, acnt, "2023-01", "")
	if err != nil {
		t.Fatal("Error creating invoice: ", err)
	}

	os.WriteFile("invoice.html", []byte(invoice), 0644)
}
