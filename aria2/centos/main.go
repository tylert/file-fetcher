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
	reg := regexp.MustCompile(`CentOS-Stream-\d+?-\d+?\.\d+`)
	ver := ""
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, ".iso") && !strings.Contains(href, "-latest") && !strings.Contains(href, ".torrent") && !strings.Contains(href, "MD5SUM") && !strings.Contains(href, "SHA1SUM") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=CentOS")
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
				}
			} else if !strings.Contains(href, "SHA256SUM") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
	// Now that we know the release number, we can give the generic checksum files sensible names
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if (strings.Contains(href, "SHA256SUM") || strings.Contains(href, "CHECKSUM")) && !strings.Contains(href, ".iso") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=CentOS")
				fmt.Println(fmt.Sprintf("	out=%s-%s", ver, href))
			}
		}
	})
}

func dumpUgh(url string) {
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

	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if !strings.Contains(href, "-latest") && !strings.Contains(href, ".torrent") && !strings.Contains(href, "MD5") && !strings.Contains(href, "SHA1SUM") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=CentOS")
			} else {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
}

func doIt() {
	fmt.Println("# https://mirror.xenyth.net/centos")
	fmt.Println("# https://centos.org/download")
	fmt.Println("# https://centos.org")
	fmt.Println("# https://en.wikipedia.org/wiki/CentOS")
	fmt.Println("# https://distrowatch.com/centos")

	// XXX FIXME TODO  Try to determine the list of "current" releases programatically!!!
	dumpOne("https://mirror.xenyth.net/centos-stream/9-stream/BaseOS/x86_64/iso")
	dumpOne("https://mirror.xenyth.net/centos/8-stream/isos/x86_64")
	dumpUgh("https://mirror.xenyth.net/centos/7/isos/x86_64")

	// Get the signing keys too
	fmt.Println("https://mirror.xenyth.net/centos/RPM-GPG-KEY-CentOS-7") // 6341AB2753D78A78A7C27BB124C6A8A7F4A80EB5
	fmt.Println("	dir=CentOS")
}

func main() {
	doIt()
}
