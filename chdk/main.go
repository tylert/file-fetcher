package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func dumpOne(url string) {
	res, err := http.Get(url)
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

	// XXX FIXME TODO  Stop after finding the first instance!!!
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if strings.Contains(href, "_elph115-") && strings.Contains(href, "full.zip") {
				fmt.Println(fmt.Sprintf("%s%s", url, href))
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	dir=CHDK")
			}
		}
	})
}

func main() {
	dumpOne("https://mighty-hoernsche.de/")
	dumpOne("https://mighty-hoernsche.de/trunk/")
}
