package cli

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Print in screen Go version and arch.
func Current() {
	goVersion, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatalf("error at getting Go version: %v", err)
	}

	output := strings.Split(string(goVersion), " ")

	fmt.Printf("%s %s", output[2], output[3])
}
