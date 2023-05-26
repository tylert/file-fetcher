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

	fmt.Println("# https://nncp.mirrors.quux.org/Tarballs.html")
	fmt.Println("# https://www.complete.org/nncp")

	// Stop after showing exactly 3 download links (which should be the newest ones)
	count := 3
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if strings.Contains(href, "download") {
				fmt.Println(fmt.Sprintf("https://nncp.mirrors.quux.org/%s", href))
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	dir=NNCP")

				count--
				if count <= 0 {
					os.Exit(0)
				}
			}
		}
	})
}
