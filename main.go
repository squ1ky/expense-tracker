package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: expense-tracker <command> [options]")
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "delete":
		handleDelete()
	case "summary":
		handleSummary()
	default:
		log.Fatalf("unknown command: %s", command)
	}
}
