package main

import (
	"os"
	"strings"
)

func parseParams() map[string]string {
	params := make(map[string]string)
	args := os.Args[2:]

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--") {
			if strings.Contains(arg, "=") {
				parts := strings.SplitN(arg, "=", 2)
				params[parts[0]] = parts[1]
			} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				params[arg] = args[i+1]
				i++
			}
		}
	}

	return params
}
