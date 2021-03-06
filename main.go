package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	var sheet string
	var column string
	var authToken string
	var skipRows int
	var tabIndex int

	flag.StringVar(&sheet, "sheet", "", "Google Sheets spreadsheet ID (Required)")
	flag.StringVar(&column, "column", "", "Column to scrape. Make sure data is in the format username/reponame (Required)")
	flag.StringVar(&authToken, "token", "", "GitHub Personal Access Token (Create one at https://github.com/settings/tokens/new) with full control of private repositories (Required)")
	flag.IntVar(&skipRows, "skip", 0, "Skip a number of rows to accommodate headers")
	flag.IntVar(&tabIndex, "tab", 0, "Spreadsheet tab to look for the specified column")

	flag.Parse()

	showBanner()

	if sheet == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if column == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if authToken == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	MakeClones(sheet, tabIndex, column, authToken, skipRows)

	// display the number of repos and length of time in the terminal.
	since := time.Since(start).Seconds()
	minutes := int(since / 60.0)
	seconds := int(since) % 60

	size := dirSize("github.com/")
	sizeStr := fmt.Sprintf("%.2f MB", size)

	if size > 1000 {
		size = size / 1000
		sizeStr = fmt.Sprintf("%.2f GB", size)
	}
	completed("Cloned %d repos in %d minutes and %d seconds", numOfReposCloned, minutes, seconds)
	pathComplete("github.com/https:/github.com/ (%s)", sizeStr)
}
