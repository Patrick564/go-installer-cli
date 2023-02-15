package main

import (
	"os"

	"github.com/Patrick564/update-go-script/internal/cli"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		cli.Default()
		os.Exit(1)
	}

	switch args[1] {
	case "help":
		cli.Default()
	case "current":
		cli.Current()
	case "list-remote":
		cli.ListRemote()
	case "install":
		cli.Install(args[2])
	default:
		cli.Help()
	}
}
