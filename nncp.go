package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	// "github.com/Masterminds/semver"
)

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		if strings.Contains(href, "download") && !strings.Contains(href, "meta4") {
			fmt.Println(href)
		}
	}
}

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

	doc.Find("a").Each(processElement)
}
