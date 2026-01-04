/*usr/bin/env go run "$0" "$@"; exit;*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	Golang()
}

func Golang() {
	// Spit out some handy links
	fmt.Println("# https://go.dev/dl")
	fmt.Println("# https://go.dev")
	fmt.Println("# https://go.dev/doc/devel/release")
	fmt.Println("# https://en.wikipedia.org/wiki/Go_(programming_language)")

	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://go.dev/dl")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Stop after showing links for exactly 1 release (which should be the newest ones)
	count := 3
	doc.Find("tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		moo := s.Find("td.filename a")
		poo := s.Find("td tt").Text()
		href, ok := moo.Attr("href")
		if ok {
			if (strings.Contains(href, "linux") && !strings.Contains(href, "386") && !strings.Contains(href, "armv6l") && !strings.Contains(href, "s390x") && !strings.Contains(href, "ppc64le")) || strings.Contains(href, "src") {
				fmt.Println(fmt.Sprintf("https://go.dev%s", href))
				fmt.Println("	dir=golang")
				fmt.Println(fmt.Sprintf("	checksum=sha-256=%s", poo))
				whatzit := strings.Split(href, "/")
				fmt.Fprintln(os.Stderr, fmt.Sprintf("%s  %s", poo, whatzit[2]))
				count--
				if count <= 0 {
					return false
				}
			} else if !strings.Contains(href, "darwin") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
		return true
	})
}
