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
				fmt.Println("	dir=BIOS")
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
	dumpOne("https://www.asus.com/ca-en/supportonly/rog%20strix%20b550-a%20gaming/helpdesk_bios")
	dumpOne("https://www.asus.com/ca-en/supportonly/rog%20strix%20b450-f%20gaming/helpdesk_bios")
	// https://pcsupport.lenovo.com/ca/en/products/laptops-and-netbooks/thinkpad-t-series-laptops/thinkpad-t560/downloads/driver-list/component?name=BIOS%2FUEFI&id=5AC6A815-321D-440E-8833-B07A93E0428C
	// https://pcsupport.lenovo.com/ca/en/products/laptops-and-netbooks/thinkpad-t-series-laptops/thinkpad-t460/downloads/driver-list/component?name=BIOS%2FUEFI&id=5AC6A815-321D-440E-8833-B07A93E0428C
	// https://pcsupport.lenovo.com/ca/en/products/laptops-and-netbooks/thinkpad-t-series-laptops/thinkpad-t450s/downloads/driver-list/component?name=BIOS%2FUEFI&id=5AC6A815-321D-440E-8833-B07A93E0428C
}

func main() {
	doIt()
}
