package cli

import (
	"fmt"
	"os"
)

// When call with a unknow command.
func Help() {
	fmt.Println("Command unknow, type `goit` or `goit help` for see all commands avaiable.")
	os.Exit(1)
}
