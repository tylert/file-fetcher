package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://www.pistar.uk/downloads")
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

	fmt.Println("# https://www.pistar.uk/downloads")
	fmt.Println("# https://www.pistar.uk")

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, ".zip") && (strings.Contains(href, "RPi") || strings.Contains(href, "dvmega")) {
				fmt.Println(fmt.Sprintf("https://www.pistar.uk/downloads/%s", href))
				fmt.Println("	dir=Pi-Star")
			}
		}
	})
}
