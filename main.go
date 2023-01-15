package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Patrick564/update-go-script/utils"
	colly "github.com/gocolly/colly/v2"
)

const (
	allowedDomain = "go.dev"
	downloadUrl   = "https://go.dev/dl/"
)

// Show commands for use.
func defaultCmd() {
	fmt.Print("This a CLI app to detect your Go version and install a new deleting previous.\n")

	fmt.Print("\nUSAGE:\n")
	fmt.Print("  goit <command>\n")

	fmt.Print("\nCOMMANDS:\n")
	fmt.Print("  current:     Get the current installed version of Go.\n")
	fmt.Print("  list-remote: List all Go versions avaiable for your syhstem distribution.\n")
	fmt.Print("  install:     Install the indicated version fo Go.\n")

	fmt.Println()
}

func helpCmd() {
	fmt.Println("Command unknow, type `goit` or `goit help` for see all commands avaiable.")
}

// Get the current installed Go version.
func currentCmd() {
	goVersion, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatalf("error at getting Go version: %v", err)
	}

	output := strings.Split(string(goVersion), " ")

	fmt.Printf("%s %s", output[2], output[3])
}

func listRemoteCmd(c *colly.Collector) {
	links := make([]string, 0)
	replacer := strings.NewReplacer("/dl/", "", ".linux-amd64.tar.gz", "")

	c.OnHTML(".filename", func(h *colly.HTMLElement) {
		link := h.ChildAttr(".download", "href")

		if strings.Contains(link, utils.SystemDist()) {
			links = append(links, link)
		}
	})
	c.Visit(downloadUrl)

	for _, l := range links[:3] {
		fmt.Println(replacer.Replace(l))
	}
}

func installCmd(version string) {
	fmt.Println(downloadUrl, version, utils.SystemDist(), utils.ExtFile())
}

func main() {
	args := os.Args
	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomain),
	)

	if len(args) == 1 || args[1] == "help" {
		defaultCmd()
		return
	}

	if args[1] == "current" {
		currentCmd()
		return
	}

	if args[1] == "list-remote" {
		listRemoteCmd(c)
		return
	}

	if args[1] == "install" {
		installCmd(args[2])
		return
	}

	helpCmd()
}
