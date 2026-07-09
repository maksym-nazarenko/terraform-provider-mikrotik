package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	if err := run(); err != nil {
		log.Println(fmt.Errorf("application failed: %w", err))
		os.Exit(1)
	}
}

func run() error {
	appName := path.Base(os.Args[0])
	if len(os.Args) < 2 {
		usage(appName)
		return nil
	}
	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "resource":
		return subcommandResource(args)
	// case "test":
	// 	return subcommandTest(args)
	case "-h", "--help":
		usage(appName)
		return nil
	default:
		usage(appName)
		return fmt.Errorf("unknown subcommand: %s", subcommand)
	}
}
