package main

import (
	"fmt"
	"log"
	"strconv"
)

const filename = "list.txt"

func handleAdd() {
	params := parseParams()

	description, ok := params["--description"]
	if !ok {
		log.Fatal("--description is required")
	}

	amountStr, ok := params["--amount"]
	if !ok {
		log.Fatal("--amount is required")
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		log.Fatalf("amount must be an integer: %v", err)
	}

	tracker := NewExpenseTracker(filename)
	expense, err := tracker.AddExpense(description, amount)
	if err != nil {
		log.Fatalf("error adding expense: %v", err)
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
}

func handleList() {
	tracker := NewExpenseTracker(filename)
	if err := tracker.PrintExpenses(); err != nil {
		log.Fatalf("error listing expenses: %v", err)
	}
}

func handleDelete() {
	params := parseParams()

	idStr, ok := params["--id"]
	if !ok {
		log.Fatal("--id is required")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("id must be an integer: %v", err)
	}

	tracker := NewExpenseTracker(filename)
	if err := tracker.DeleteExpense(id); err != nil {
		log.Fatalf("error deleting expense: %v", err)
	}

	fmt.Printf("Expense with ID %d deleted successfully\n", id)
}

func handleSummary() {
	tracker := NewExpenseTracker(filename)
	total, err := tracker.GetSummary()
	if err != nil {
		log.Fatalf("error getting summary: %v", err)
	}

	fmt.Printf("Total expenses: $%d\n", total)
}
