package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: command [URL]")
		os.Exit(1)
	}
	start := os.Args[1]
	doc, err := goquery.NewDocument(start)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	csvfile, _ := os.Create("apple-api-diff.csv")
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	writer.Write([]string{"Status", "Framework", "Target"})
	doc.Find(".diffReport2 a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href") // Always exists
		startURL, _ := url.Parse(start)
		hrefURL, _ := url.Parse(href)
		doc, err := goquery.NewDocument(startURL.ResolveReference(hrefURL).String())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		h2 := doc.Find(".diffReport2 h2").Text()
		doc.Find(".symbolName").Each(func(i int, s *goquery.Selection) {
			status := s.Find(".diffStatus").Text()
			target := s.Text()
			writer.Write([]string{status, h2, target})
		})
	})
	writer.Flush()
}
