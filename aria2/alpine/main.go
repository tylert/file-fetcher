package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func dumpOne(url string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(url)
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

	fmt.Println(fmt.Sprintf("# %s", url))

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.HasPrefix(href, "alpine-standard-") || strings.HasPrefix(href, "alpine-rpi-") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
			} else {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
}

func doIt() {
	fmt.Println("# https://dl-cdn.alpinelinux.org/alpine/latest-stable")
	fmt.Println("# https://alpinelinux.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Alpine_Linux")
	fmt.Println("# https://distrowatch.com/table.php?distribution=alpine")

	dumpOne("https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/x86_64")
	dumpOne("https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/aarch64")
}

func main() {
	doIt()
}
