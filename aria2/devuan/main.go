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
	Devuan()
}

func Devuan() {
	// Spit out some handy links
	fmt.Println("# https://mirror.leaseweb.com/devuan")
	fmt.Println("# https://devuan.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Devuan")
	fmt.Println("# https://distrowatch.com/devuan")

	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	// XXX FIXME TODO  Hunt for files that are named LATEST_STABLE_IS_* and convert it to lowercase
	// https://mirror.leaseweb.com/devuan/LATEST_STABLE_IS_EXCALIBUR
	res, err := client.Get("https://mirror.leaseweb.com/devuan/devuan_excalibur/minimal-live")
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

	// Moo
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "devuan") {
				fmt.Println(fmt.Sprintf("https://mirror.leaseweb.com/devuan/devuan_excalibur/minimal-live/%s", href))
				fmt.Println("	dir=Linux/Devuan")
			}
		}
	})
}
