package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          int
	Date        time.Time
	Description string
	Amount      int
}

func (e Expense) String() string {
	return fmt.Sprintf("%d %s %s %d", e.Id, e.Date.Format("2006-01-02"), e.Description, e.Amount)
}
