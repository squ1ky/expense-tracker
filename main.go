package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.OpenFile("list.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = fmt.Errorf("error closing file: %v", err)
		}
	}(file)

	args := os.Args
	params := parseParams()

	switch args[1] {
	case "add":
		amount, err := strconv.Atoi(params["--amount"])

		if err != nil {
			_ = fmt.Errorf("amount must be an integer")
		}

		id, err := MaxId(file)

		if err != nil || id == -1 {
			_ = fmt.Errorf("error getting max id: %v", err)
		}

		expense := Expense{
			Id:          id + 1,
			Date:        time.Now(),
			Description: params["--description"],
			Amount:      amount,
		}

		err = AddExpense(expense, file)
		if err != nil {
			_ = fmt.Errorf("error adding expense")
		} else {
			fmt.Printf("Expense added successfully (ID: %d)", expense.Id)
		}
	case "list":
		PrintExpenses(file)
	case "delete":
		id, err := strconv.Atoi(params["--id"])
		if err != nil {
			_ = fmt.Errorf("id must be an integer")
		}
		err = DeleteExpense(id, file)
		if err != nil {
			_ = fmt.Errorf("error deleting expense")
		}
	case "summary":
		total, err := Summary(file)
		if err != nil {
			_ = fmt.Errorf("error getting summary: %v", err)
		}
		fmt.Printf("Total expenses: $%d\n", total)
	}
}

func parseParams() map[string]string {
	params := make(map[string]string)

	for i, arg := range os.Args[1:] {
		realI := i + 1

		if strings.HasPrefix(arg, "--") && strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			params[parts[0]] = parts[1]
		}

		if strings.HasPrefix(arg, "--") && !strings.Contains(arg, "=") {
			if realI+1 < len(os.Args) && !strings.Contains(os.Args[realI+1], "--") {
				params[arg] = os.Args[realI+1]
			}
		}
	}

	return params
}
