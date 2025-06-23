package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type ExpenseTracker struct {
	filename string
}

func NewExpenseTracker(filename string) *ExpenseTracker {
	return &ExpenseTracker{filename: filename}
}

func (et *ExpenseTracker) AddExpense(description string, amount int) (*Expense, error) {
	maxID, err := et.getMaxID()
	if err != nil {
		return nil, fmt.Errorf("failed to get max ID: %w", err)
	}

	expense := &Expense{
		ID:          maxID + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}

	file, err := os.OpenFile(et.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(expense.String() + "\n"); err != nil {
		return nil, fmt.Errorf("failed to write expense: %w", err)
	}

	return expense, nil
}

func (et *ExpenseTracker) PrintExpenses() error {
	expenses, err := et.loadExpenses()
	if err != nil {
		return fmt.Errorf("failed to load expenses: %w", err)
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}

	fmt.Printf("%-5s %-12s %-20s %-10s\n", "ID", "Date", "Description", "Amount")
	fmt.Println(strings.Repeat("-", 50))

	for _, expense := range expenses {
		fmt.Printf("%-5d %-12s %-20s $%-10d\n",
			expense.ID,
			expense.Date.Format("2006-01-02"),
			expense.Description,
			expense.Amount)
	}

	return nil
}

func (et *ExpenseTracker) DeleteExpense(id int) error {
	expenses, err := et.loadExpenses()
	if err != nil {
		return fmt.Errorf("failed to load expenses: %w", err)
	}

	var filteredExpenses []*Expense
	found := false

	for _, expense := range expenses {
		if expense.ID != id {
			filteredExpenses = append(filteredExpenses, expense)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	return et.saveExpenses(filteredExpenses)
}

func (et *ExpenseTracker) GetSummary() (int, error) {
	expenses, err := et.loadExpenses()
	if err != nil {
		return 0, fmt.Errorf("failed to load expenses: %w", err)
	}

	total := 0
	for _, expense := range expenses {
		total += expense.Amount
	}

	return total, nil
}

func (et *ExpenseTracker) getMaxID() (int, error) {
	expenses, err := et.loadExpenses()
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, expense := range expenses {
		if expense.ID > maxID {
			maxID = expense.ID
		}
	}

	return maxID, nil
}

func (et *ExpenseTracker) loadExpenses() ([]*Expense, error) {
	file, err := os.Open(et.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []*Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var expenses []*Expense
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		expense, err := ParseExpense(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse expense: %w", err)
		}

		expenses = append(expenses, expense)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return expenses, nil
}

func (et *ExpenseTracker) saveExpenses(expenses []*Expense) error {
	file, err := os.Create(et.filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	for _, expense := range expenses {
		if _, err := file.WriteString(expense.String() + "\n"); err != nil {
			return fmt.Errorf("failed to write expense: %w", err)
		}
	}

	return nil
}
