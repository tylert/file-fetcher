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

	fmt.Println(fmt.Sprintf("# %s", url))

	// Stop after showing links for exactly 1 release (which should be the newest ones)
	count := 1
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, ".zip") {
				moo := href[:strings.IndexByte(href, '?')]
				fmt.Println(moo)
				fmt.Println("	dir=bios")
				count--
				if count <= 0 {
					return false
				}
			}
		}
		return true
	})
}

func doIt() {
	dumpOne("https://www.asus.com/supportonly/rog%20strix%20b550-a%20gaming/helpdesk_bios")
	dumpOne("https://www.asus.com/supportonly/rog%20strix%20b450-f%20gaming/helpdesk_bios")
}

func main() {
	doIt()
}
