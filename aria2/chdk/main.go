package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func dumpOne(url string, target string) {
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

	fmt.Println(fmt.Sprintf("# %s", url))

	// Break after we find a single match (e.g.:  device twins)
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			// Don't bother with the partial builds, only full ones
			if strings.Contains(href, target) && strings.Contains(href, "full.zip") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=CHDK")
				return false
			}
		}
		return true
	})
}

func doIt() {
	fmt.Println("# https://chdk.fandom.com/wiki/CHDK")
	fmt.Println("# https://en.wikipedia.org/wiki/DIGIC#CHDK")

	dumpOne("https://mighty-hoernsche.de", "_elph115-")       // stable
	dumpOne("https://mighty-hoernsche.de/trunk", "_elph115-") // unstable
}

func main() {
	doIt()
}
