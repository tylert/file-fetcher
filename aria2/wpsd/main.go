package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func doIt() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://w0chp.radio/wpsd")
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
	fmt.Println("# https://w0chp.radio/wpsd")
	fmt.Println("# https://manual.wpsd.radio")

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "RPi") || strings.Contains(href, "SHA256") {
				fmt.Println(fmt.Sprintf("%s", href))
				fmt.Println("	dir=WPSD")
			} else if strings.Contains(href, "WPSD") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
}

func main() {
	doIt()
}
