package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func doIt() {
	res, err := http.Get("https://nncp.mirrors.quux.org/Tarballs.html")
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

	fmt.Println("# https://nncp.mirrors.quux.org/Tarballs.html")
	fmt.Println("# https://www.complete.org/nncp")

	// Stop after showing links for exactly 1 release (which should be the newest ones)
	count := 2
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "download") && !strings.Contains(href, ".meta4") {
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
