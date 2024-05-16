package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
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

	// Do a first pass to get the version number to use when renaming the checksum files
	reg1 := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	reg2 := regexp.MustCompile(`\d+?\.\d+`)
	ver := ""
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if (strings.Contains(href, ".iso") || strings.Contains(href, ".zsync") || strings.Contains(href, ".list") || strings.Contains(href, ".manifest")) && !strings.Contains(href, ".torrent") && !strings.Contains(href, "-desktop") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Ubuntu")
				if reg1.FindString(href) != "" { // e.g.:  22.04.3
					ver = reg1.FindString(href)
				} else if reg2.FindString(href) != "" { // e.g.:  23.04
					ver = reg2.FindString(href)
				}
			} else if !strings.Contains(href, "SHA256SUMS") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
	// Now that we know the release number, we can give the generic checksum files sensible names
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "SHA256SUMS") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Ubuntu")
				fmt.Println(fmt.Sprintf("	out=ubuntu-%s-%s.txt", ver, href))
			}
		}
	})
}

func doIt() {
	// Spit out some handy links
	fmt.Println("# https://mirror.xenyth.net/ubuntu-releases")
	fmt.Println("# https://releases.ubuntu.com")
	fmt.Println("# https://cdimage.ubuntu.com")
	fmt.Println("# https://en.wikipedia.org/wiki/Ubuntu_version_history#Table_of_versions")
	fmt.Println("# https://en.wikipedia.org/wiki/Ubuntu")
	fmt.Println("# https://distrowatch.com/ubuntu")

	// XXX FIXME TODO  Try to determine the list of "current" releases programatically!!!
	dumpOne("https://mirror.xenyth.net/ubuntu-releases/24.04") // noble (until 2029-05-31)
	dumpOne("https://mirror.xenyth.net/ubuntu-releases/22.04") // jammy (until 2027-06-01)
	dumpOne("https://mirror.xenyth.net/ubuntu-releases/20.04") // focal (until 2025-05-29)

	// Spit out some more handy links
	fmt.Println("# https://cloud-images.ubuntu.com")
	fmt.Println("# https://cloud-images.ubuntu.com/locator")
	fmt.Println("# https://ubuntu.com/download/raspberry-pi")
}

func main() {
	doIt()
}
