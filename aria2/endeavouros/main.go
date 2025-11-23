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
	res, err := client.Get("https://endeavouros.com/#Download")
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
	fmt.Println("# https://endeavouros.com")
	fmt.Println("# https://endeavouros-team.github.io/EndeavourOS-Development")
	fmt.Println("# https://github.com/endeavouros-team/ISO/releases/latest")
	fmt.Println("# https://en.wikipedia.org/wiki/EndeavourOS")
	fmt.Println("# https://distrowatch.com/endeavour")

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "gigenet") && strings.Contains(href, ".iso") && !strings.Contains(href, ".torrent") {
				fmt.Println(fmt.Sprintf("%s", href))
				fmt.Println("	dir=Linux/EndeavourOS")
			}
		}
	})
}

func main() {
	doIt()
}
