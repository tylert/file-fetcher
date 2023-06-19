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
	reg := regexp.MustCompile(`\d+?\.\d+`)
	ver := ""
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if !strings.Contains(href, "Rocky-x86_64") && strings.Contains(href, ".iso") && !strings.Contains(href, ".CHECKSUM") && !strings.Contains(href, ".torrent") && !strings.Contains(href, "-latest") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Rocky")
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
				}
			} else if !strings.Contains(href, "CHECKSUM") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
	// Now that we know the release number, we can give the generic checksum files sensible names
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if !strings.Contains(href, "Rocky-x86_64") && strings.Contains(href, "CHECKSUM") && !strings.Contains(href, ".CHECKSUM") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Rocky")
				fmt.Println(fmt.Sprintf("	out=Rocky-%s-x86_64-%s", ver, href))
			}
		}
	})
}

func main() {
	fmt.Println("# https://mirror.xenyth.net/rocky")
	fmt.Println("# https://rockylinux.org/download")
	fmt.Println("# https://rockylinux.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Rocky_Linux")
	fmt.Println("# https://distrowatch.com/rocky")

	// XXX FIXME TODO  Try to determine the list of "current" releases programatically!!!
	dumpOne("https://mirror.xenyth.net/rocky/9/isos/x86_64")
	dumpOne("https://mirror.xenyth.net/rocky/8/isos/x86_64")

	// Grab the signing keys too
	fmt.Println("https://mirror.xenyth.net/rocky/RPM-GPG-KEY-Rocky-9") // 21CB256AE16FC54C6E652949702D426D350D275D
	fmt.Println("	dir=Rocky")
	fmt.Println("https://mirror.xenyth.net/rocky/RPM-GPG-KEY-Rocky-8") // 7051C470A929F454CEBE37B715AF5DAC6D745A60
	fmt.Println("	dir=Rocky")
}
