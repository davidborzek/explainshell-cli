package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func explain(cmd string) {
	res, err := http.Get(
		fmt.Sprintf("https://explainshell.com/explain?cmd=%s", url.QueryEscape(cmd)),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	delimiter := strings.Repeat("_", 50)
	doc.Find(".help-box").Each(func(_ int, s *goquery.Selection) {
		fmt.Printf("%s\n%s\n\n", s.Text(), delimiter)
	})
}

func main() {
	command := strings.Join(os.Args[1:], " ")
	if command == "" {
		fmt.Fprintf(os.Stderr, "usage: %s [command]\n", os.Args[0])
		os.Exit(1)
	}

	explain(command)
}
