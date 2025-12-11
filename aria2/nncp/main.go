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
	res, err := client.Get("https://nncp.mirrors.quux.org/Tarballs.html")
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
	fmt.Println("# https://nncp.mirrors.quux.org/Tarballs.html")
	fmt.Println("# https://nncp.mirrors.quux.org")
	fmt.Println("# https://complete.org/nncp")
	fmt.Println("# http://www.nncpgo.org")

	// Newest releases are at the top / Single table row per release
	count := 4
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "download") {
				fmt.Println(fmt.Sprintf("https://nncp.mirrors.quux.org/%s", href))
				fmt.Println("	dir=NNCP")
				count--
				if count <= 0 {
					return false
				}
			}
		}
		return true
	})
}

func main() {
	doIt()
}
