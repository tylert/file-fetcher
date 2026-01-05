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
	Alpine()
}

func Alpine() {
	fmt.Println("# https://dl-cdn.alpinelinux.org/alpine/latest-stable")
	fmt.Println("# https://alpinelinux.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Alpine_Linux")
	fmt.Println("# https://distrowatch.com/alpine")

	dumpOne("https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/x86_64")
	dumpOne("https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/aarch64")
}

func dumpOne(url string) {
	fmt.Println(fmt.Sprintf("# %s", url))

	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(url)
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

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.HasPrefix(href, "alpine-standard-") && strings.Contains(href, "-x86_64.iso") && !strings.Contains(href, "rc1") && !strings.Contains(href, "rc2") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Linux/Alpine")
			} else if strings.HasPrefix(href, "alpine-rpi-") && strings.Contains(href, "-aarch64.img.gz") && !strings.Contains(href, "rc1") && !strings.Contains(href, "rc2") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Linux/Alpine")
			} else {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
}
