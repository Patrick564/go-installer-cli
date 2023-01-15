package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	colly "github.com/gocolly/colly/v2"
)

// goPath, err := exec.Command("go", "env", "GOPATH").Output()
// if err != nil {
// 	log.Fatalf("error at getting Go path: %v", err)
// }

// Show commands for use.
func defaultCmd() {
	fmt.Print("This a CLI app to detect your Go version and install a new deleting previous.\n")

	fmt.Print("\nUSAGE:\n")
	fmt.Print("  goit <command>\n")

	fmt.Print("\nCOMMANDS:\n")
	fmt.Print("  current:     Get the current installed version of Go\n")
	fmt.Print("  list-remote: Get the current installed version of Go\n")
	fmt.Print("  install:     Get the current installed version of Go\n")

	fmt.Println()
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
	system := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	replacer := strings.NewReplacer("/dl/", "", ".linux-amd64.tar.gz", "")

	c.OnHTML(".filename", func(h *colly.HTMLElement) {
		link := h.ChildAttr(".download", "href")

		if strings.Contains(link, system) {
			links = append(links, link)
		}
	})
	c.Visit("https://go.dev/dl/")

	for _, l := range links[:3] {
		fmt.Println(replacer.Replace(l))
	}
}

func installCmd(version string) {
	fmt.Println(version)
}

func main() {
	args := os.Args
	c := colly.NewCollector(
		colly.AllowedDomains("go.dev"),
	)

	if len(args) == 1 {
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
}
