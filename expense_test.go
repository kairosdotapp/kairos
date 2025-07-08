package main

import (
	"testing"
	"time"
)

func TestGetExpensesForCustAIn202506(t *testing.T) {
	expenseTracker, err := NewExpenseTracker("example/expenses.csv")
	if err != nil {
		t.Fatalf("NewExpenseTracker failed: %v", err)
	}

	startDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 6, 30, 23, 59, 59, 0, time.UTC)

	expenses, err := expenseTracker.GetExpensesForTimeRangeAndCustomer(startDate, endDate, "CustA")
	if err != nil {
		t.Fatalf("GetExpensesForTimeRangeAndCustomer failed: %v", err)
	}

	expectedCount := 1
	if len(expenses) != expectedCount {
		t.Errorf("Expected %d expenses for CustA in 2025-06, got %d", expectedCount, len(expenses))
	}

	if len(expenses) > 0 {
		expense := expenses[0]
		if expense.Customer != "CustA" {
			t.Errorf("Expected customer 'CustA', got '%s'", expense.Customer)
		}
		if expense.Amount != 131.11 {
			t.Errorf("Expected amount 131.11, got %.2f", expense.Amount)
		}
		if expense.Vendor != "Digikey" {
			t.Errorf("Expected vendor 'Digikey', got '%s'", expense.Vendor)
		}
	}
}

func TestGetTotalExpensesForCustAFrom20250624To20250702(t *testing.T) {
	expenseTracker, err := NewExpenseTracker("example/expenses.csv")
	if err != nil {
		t.Fatalf("NewExpenseTracker failed: %v", err)
	}

	startDate := time.Date(2025, 6, 24, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 7, 2, 23, 59, 59, 0, time.UTC)

	expenses, err := expenseTracker.GetExpensesForTimeRangeAndCustomer(startDate, endDate, "CustB")
	if err != nil {
		t.Fatalf("GetExpensesForTimeRangeAndCustomer failed: %v", err)
	}

	var totalAmount float64
	for _, expense := range expenses {
		totalAmount += expense.Amount
	}

	expectedTotal := 524.69
	tolerance := 0.01
	if totalAmount < expectedTotal-tolerance || totalAmount > expectedTotal+tolerance {
		t.Errorf("Expected total amount %.2f for CustB from 2025-06-24 to 2025-07-02, got %.2f", expectedTotal, totalAmount)
	}

	expectedCount := 4
	if len(expenses) != expectedCount {
		t.Errorf("Expected %d expenses for CustB in date range, got %d", expectedCount, len(expenses))
	}
}
