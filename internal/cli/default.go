package cli

import "fmt"

// Show all commands aviable for the CLI.
func Default() {
	fmt.Print("This a CLI app to detect your Go version and install a new deleting previous.\n")

	fmt.Print("\nUSAGE:\n")
	fmt.Print("  goit <command>\n")

	fmt.Print("\nCOMMANDS:\n")
	fmt.Print("  current:     Get the current installed version of Go.\n")
	fmt.Print("  list-remote: List all Go versions avaiable for your syhstem distribution.\n")
	fmt.Print("  install:     Install the indicated version fo Go.\n")

	fmt.Println()
}
