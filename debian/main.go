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

	// XXX FIXME TODO  Figure out why each link gets printed twice!!!
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if strings.Contains(href, "netinst.iso") && !strings.Contains(href, "-edu-") && !strings.Contains(href, "-mac-") {
				fmt.Println(fmt.Sprintf("%s%s", url, href))
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	conditional-get=true")
			}
			if strings.Contains(href, "SHA") {
				fmt.Println(fmt.Sprintf("%s%s.txt", url, href))
				fmt.Println("	auto-file-renaming=false")
				// fmt.Println("	out=moo")
			}
		}
	})
}

func main() {
	const (
		testing      = "http://cdimage.debian.org/cdimage/weekly-builds/amd64/iso-cd/"
		stable       = "http://cdimage.debian.org/cdimage/release/current/amd64/iso-cd/"
		oldstable    = "http://cdimage.debian.org/cdimage/archive/latest-oldstable/amd64/iso-cd/"
		oldoldstable = "http://cdimage.debian.org/cdimage/archive/latest-oldoldstable/amd64/iso-cd/"
	)

	dumpOne(testing)
	dumpOne(stable)
	dumpOne(oldstable)
	dumpOne(oldoldstable)
}
