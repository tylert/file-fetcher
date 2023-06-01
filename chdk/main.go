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

	doc.Find("a").EachWithBreak(func(index int, element *goquery.Selection) bool {
		href, exists := element.Attr("href")
		if exists {
			if strings.Contains(href, "_elph115-") && strings.Contains(href, "full.zip") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	dir=CHDK")
				fmt.Println("	file-allocation=falloc")
				return false
			}
		}
		return true
	})
}

func main() {
	fmt.Println("# https://mighty-hoernsche.de")
	fmt.Println("# https://mighty-hoernsche.de/trunk")
	fmt.Println("# https://chdk.fandom.com/wiki/CHDK")
	fmt.Println("# https://en.wikipedia.org/wiki/DIGIC#CHDK")

	dumpOne("https://mighty-hoernsche.de")
	dumpOne("https://mighty-hoernsche.de/trunk")
}
