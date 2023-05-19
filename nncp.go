package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
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

	count := 0
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if strings.Contains(href, "download") && !strings.Contains(href, "meta4") {
				fmt.Println(fmt.Sprintf("https://nncp.mirrors.quux.org/%s", href))
				count++
				if count > 1 {
					os.Exit(0)
				}
			}
		}
	})
}
