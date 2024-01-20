package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func doIt() {
	res, err := http.Get("https://pistar.uk/downloads")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body.", err)
	}

	// Spit out some handy links
	fmt.Println("# https://pistar.uk/downloads")
	fmt.Println("# https://pistar.uk")

	// FIXME XXX TODO  Don't include older releases here!!!
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, ".zip") && (strings.Contains(href, "RPi") || strings.Contains(href, "dvmega")) {
				fmt.Println(fmt.Sprintf("https://pistar.uk/downloads/%s", href))
				fmt.Println("	dir=Pi-Star")
			}
		}
	})
}

func main() {
	doIt()
}
