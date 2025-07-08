package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type CSVExpense struct {
	Date        string  `csv:"Date"`
	Customer    string  `csv:"Customer"`
	Vendor      string  `csv:"Vendor"`
	Description string  `csv:"Description"`
	Amount      float64 `csv:"Amount"`
}

type Expense struct {
	Date        time.Time
	Customer    string
	Vendor      string
	Description string
	Amount      float64
}

type ExpenseTracker struct {
	expenses []Expense
}

func NewExpenseTracker(csvFilePath string) (*ExpenseTracker, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	var csvExpenses []CSVExpense
	if err := gocsv.UnmarshalFile(file, &csvExpenses); err != nil {
		return nil, fmt.Errorf("failed to parse CSV file: %w", err)
	}

	var expenses []Expense
	for _, csvExpense := range csvExpenses {
		date, err := time.Parse("2006-01-02", csvExpense.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date '%s': %w", csvExpense.Date, err)
		}

		expense := Expense{
			Date:        date,
			Customer:    csvExpense.Customer,
			Vendor:      csvExpense.Vendor,
			Description: csvExpense.Description,
			Amount:      csvExpense.Amount,
		}
		expenses = append(expenses, expense)
	}

	return &ExpenseTracker{expenses: expenses}, nil
}

func (et *ExpenseTracker) GetExpensesForTimeRangeAndCustomer(startDate, endDate time.Time, customer string) ([]Expense, error) {
	var filteredExpenses []Expense

	for _, expense := range et.expenses {
		if expense.Customer == customer && 
		   (expense.Date.Equal(startDate) || expense.Date.After(startDate)) &&
		   (expense.Date.Equal(endDate) || expense.Date.Before(endDate)) {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}

	return filteredExpenses, nil
}