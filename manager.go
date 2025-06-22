package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddExpense(e Expense, file *os.File) error {
	_, err := file.WriteString(e.String() + "\n")

	if err != nil {
		return err
	}

	return nil
}

func DeleteExpense(id int, file *os.File) error {
	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		expenseInfo := strings.SplitN(line, " ", 4)
		curId, err := strconv.Atoi(expenseInfo[0])

		if err != nil {
			return err
		}

		if curId != id {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	content := strings.Join(lines, "\n")
	if len(content) > 0 {
		content += "\n"
	}

	return os.WriteFile(file.Name(), []byte(content), 0644)
}

func PrintExpenses(file *os.File) {
	fmt.Printf("%-5s %-10s %-10s %-6s\n", "ID", "Date", "Description", "Amount")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		expenseInfo := strings.SplitN(line, " ", 4)
		fmt.Printf("%-5s %-10s %-10s $%-6s\n", expenseInfo[0], expenseInfo[1], expenseInfo[2], expenseInfo[3])
	}
}

func MaxId(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	maxId := -1

	for scanner.Scan() {
		line := scanner.Text()
		expenseInfo := strings.SplitN(line, " ", 4)
		id, err := strconv.Atoi(expenseInfo[0])
		if err != nil {
			return -1, err
		}

		if id > maxId {
			maxId = id
		}
	}

	return maxId, nil
}

func Summary(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		expenseInfo := strings.SplitN(line, " ", 4)
		amount, err := strconv.Atoi(expenseInfo[3])
		if err != nil {
			return 0, err
		}
		total += amount
	}

	return total, nil
}
