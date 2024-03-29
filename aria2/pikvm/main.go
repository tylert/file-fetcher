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
	res, err := http.Get("https://files.pikvm.org/images")
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
	fmt.Println("# https://files.pikvm.org/images")
	fmt.Println("# https://docs.pikvm.org/flashing_os")
	fmt.Println("# https://pikvm.org/download")
	fmt.Println("# https://pikvm.org")
	fmt.Println("# https://docs.pikvm.org")
	fmt.Println("# https://github.com/pikvm/pikvm")
	fmt.Println("# https://en.wikipedia.org/wiki/Pi-KVM")
	fmt.Println("# https://kickstarter.com/projects/mdevaev/pikvm-v4")

	// Compiled binaries
	now := time.Now()
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "v2-hdmi-rpi4-latest") {
				fmt.Println(fmt.Sprintf("https://files.pikvm.org/images/%s", href))
				fmt.Println("	dir=PiKVM")
				thingy := strings.Split(strings.ReplaceAll(href, "latest", now.Format("2006-01-02")), "/")
				fmt.Println(fmt.Sprintf("	out=pikvm-%s", thingy[len(thingy)-1]))
			} else if strings.Contains(href, "v2-hdmi-rpi4-") && strings.Contains(href, ".img") {
				whatzit := strings.Split(href, "/")
				fmt.Println(fmt.Sprintf("# skipped %s", whatzit[len(whatzit)-1]))
			}
		}
	})

	// Documentation
	fmt.Println("https://github.com/pikvm/pikvm/archive/refs/heads/gh-pages.zip")
	fmt.Println("	dir=PiKVM")
	fmt.Println(fmt.Sprintf("	out=pikvm-handbook-%s.zip", now.Format("2006-01-02")))
	fmt.Println("https://github.com/pikvm/pikvm/archive/refs/heads/master.zip")
	fmt.Println("	dir=PiKVM")
	fmt.Println(fmt.Sprintf("	out=pikvm-more-docs-%s.zip", now.Format("2006-01-02")))
}

func main() {
	doIt()
}
