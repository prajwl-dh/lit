package main

import (
	"fmt"
	"os"

	"github.com/prajwl-dh/lit/cmd"
	logger "github.com/prajwl-dh/lit/util"
)

func main() {
	if len(os.Args) < 2 {
		logger.PrintInfo("Usage: %s <command>\nTo see all commands, run: %s help\n", os.Args[0], os.Args[0])
		fmt.Println()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		logger.PrintInfo("Command: ./lit init    Usage: Initialize the repository\n")
		fmt.Println()
		os.Exit(0)

	case "init":
		if err := cmd.Init(); err != nil {
			logger.PrintError("Error: %v\n", err)
			fmt.Println()
			os.Exit(1)
		}

	default:
		logger.PrintError("Unknown command: %s\n", os.Args[1])
		fmt.Println()
		os.Exit(1)
	}
}
