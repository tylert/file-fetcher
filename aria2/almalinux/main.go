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
			if strings.Contains(href, ".iso") && !strings.Contains(href, "-latest") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=AlmaLinux")
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
			if strings.Contains(href, "CHECKSUM") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=AlmaLinux")
				fmt.Println(fmt.Sprintf("	out=AlmaLinux-%s-x86_64-%s", ver, href))
			}
		}
	})
}

func doIt() {
	// Spit out some handy links
	fmt.Println("# https://mirror.xenyth.net/almalinux")
	fmt.Println("# https://mirrors.almalinux.org/isos.html")
	fmt.Println("# https://almalinux.org")
	fmt.Println("# https://en.wikipedia.org/wiki/AlmaLinux#Releases")
	fmt.Println("# https://en.wikipedia.org/wiki/AlmaLinux")
	fmt.Println("# https://distrowatch.com/alma")

	// XXX FIXME TODO  Try to determine the list of "current" releases programatically!!!
	dumpOne("https://mirror.xenyth.net/almalinux/9/isos/x86_64")
	dumpOne("https://mirror.xenyth.net/almalinux/8/isos/x86_64")
}

func main() {
	doIt()
}
