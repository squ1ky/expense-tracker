package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	ID          int
	Date        time.Time
	Description string
	Amount      int
}

func (e Expense) String() string {
	return fmt.Sprintf("%d %s %s %d",
		e.ID,
		e.Date.Format("2006-01-02"),
		e.Description,
		e.Amount)
}

func ParseExpense(line string) (*Expense, error) {
	parts := strings.SplitN(line, " ", 4)
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid expense format: %s", line)
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid expense ID: %s", parts[0])
	}

	date, err := time.Parse("2006-01-02", parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %s", parts[1])
	}

	amount, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %s", parts[3])
	}

	return &Expense{
		ID:          id,
		Date:        date,
		Description: parts[2],
		Amount:      amount,
	}, nil
}
