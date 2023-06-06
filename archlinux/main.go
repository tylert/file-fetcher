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

	// XXX FIXME TODO  Add in the bootstrap tarballs too!!!
	// Do a first pass to get the version number to use when renaming the checksum files
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := ""
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if !strings.Contains(href, "archlinux-x86_64") && strings.Contains(href, ".iso") && !strings.Contains(href, ".torrent") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	allow-overwrite=true")
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	dir=ArchLinux")
				fmt.Println("	file-allocation=falloc")
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
				}
			}
		}
	})
	// Now that we know the release number, we can give the checksum files sensible names
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "sums.txt") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	dir=ArchLinux")
				fmt.Println("	file-allocation=falloc")
				fmt.Println(fmt.Sprintf("	out=archlinux-%s-x86_64-%s", ver, href))
			}
		}
	})
}

func main() {
	fmt.Println("# https://mirror.xenyth.net/archlinux")
	fmt.Println("# https://archlinux.org/download")
	fmt.Println("# https://archlinux.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Arch_Linux")

	dumpOne("https://mirror.xenyth.net/archlinux/iso/latest")
}
