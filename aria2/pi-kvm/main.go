package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://pikvm.org/download")
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

	fmt.Println("# https://pikvm.org/download")
	fmt.Println("# https://pikvm.org")
	fmt.Println("# https://docs.pikvm.org")
	fmt.Println("# https://github.com/pikvm/pikvm")
	fmt.Println("# https://en.wikipedia.org/wiki/Pi-KVM")
	fmt.Println("# https://www.kickstarter.com/projects/mdevaev/pikvm-v4")

	// XXX FIXME TODO  Rename files with the date they were downloaded/updated!!!
	// Compiled binaries
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "v2-hdmi-rpi4-") {
				fmt.Println(href)
				fmt.Println("	dir=Pi-KVM")
				//fmt.Println(fmt.Sprintf("	out=%s", today))
			} else if strings.Contains(href, ".img") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
	// XXX FIXME TODO  Fetch all the source (docs) too!!!
}