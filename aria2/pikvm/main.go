/*usr/bin/env go run "$0" "$@"; exit;*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	PiKVM()
}

func PiKVM() {
	// Spit out some handy links
	fmt.Println("# https://files.pikvm.org/images")
	fmt.Println("# https://docs.pikvm.org/flashing_os")
	fmt.Println("# https://pikvm.org")
	fmt.Println("# https://pikvm.org/download")
	fmt.Println("# https://docs.pikvm.org")
	fmt.Println("# https://github.com/pikvm/pikvm")
	fmt.Println("# https://en.wikipedia.org/wiki/Pi-KVM")
	fmt.Println("# https://kickstarter.com/projects/mdevaev/pikvm-v4")

	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://files.pikvm.org/images/v2-hdmi-rpi4/aarch64")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { // 200
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Compiled binaries
	now := time.Now()
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "latest") {
				fmt.Println(fmt.Sprintf("https://files.pikvm.org/images/v2-hdmi-rpi4/aarch64/%s", href))
				fmt.Println("	dir=Linux/PiKVM")
				thingy := strings.Split(strings.ReplaceAll(href, "latest", now.Format("2006-01-02")), "/")
				fmt.Println(fmt.Sprintf("	out=pikvm-%s", thingy[len(thingy)-1]))
			} else if strings.Contains(href, "v2-hdmi-rpi4-aarch64-") && strings.Contains(href, ".img") {
				whatzit := strings.Split(href, "/")
				fmt.Println(fmt.Sprintf("# skipped %s", whatzit[len(whatzit)-1]))
			}
		}
	})

	// Documentation
	fmt.Println("https://github.com/pikvm/pikvm/archive/refs/heads/gh-pages.zip")
	fmt.Println("	dir=Linux/PiKVM")
	fmt.Println(fmt.Sprintf("	out=pikvm-handbook-%s.zip", now.Format("2006-01-02")))
	fmt.Println("https://github.com/pikvm/pikvm/archive/refs/heads/master.zip")
	fmt.Println("	dir=Linux/PiKVM")
	fmt.Println(fmt.Sprintf("	out=pikvm-more-docs-%s.zip", now.Format("2006-01-02")))
}
