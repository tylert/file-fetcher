/*usr/bin/env go run "$0" "$@"; exit;*/
package main

import (
	"fmt"
	//"log"
	//"net/http"
	//"strings"
	//"github.com/PuerkitoBio/goquery"
)

func main() {
	// XXX FIXME TODO  Sigh, the CHDK builds page is all Jabbascript... https://github.com/robertkrimen/otto
	// https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks#handle-javascript-based-pages
	// https://gist.github.com/cryptix/87127f76a94183747b53

	/*
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		res, err := client.Get("https://build.chdk.photos/#build=release/IXUS/ixus132_elph115/100b")
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
	*/

	// Spit out some handy links
	fmt.Println("# https://en.wikipedia.org/wiki/DIGIC#CHDK")
	fmt.Println("# https://chdk.fandom.com/wiki/Downloads")
	fmt.Println("# https://build.chdk.photos/#build=trunk/IXUS/ixus132_elph115/100b")
	fmt.Println("# https://build.chdk.photos/#build=release/IXUS/ixus132_elph115/100b")
	fmt.Println("# https://build.chdk.photos")
	fmt.Println("# https://app.assembla.com/spaces/chdk/subversion/source")

	/*
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if ok {
				if strings.Contains(href, "full.zip") {
					fmt.Println(fmt.Sprintf("%s", href))
					fmt.Println("	dir=CHDK")
				}
			}
		})
	*/
}
