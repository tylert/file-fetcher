/*usr/bin/env go run "$0" "$@"; exit;*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	Arch()
}

func Arch() {
	// Spit out some handy links
	fmt.Println("# https://mirror.xenyth.net/archlinux")
	fmt.Println("# https://archlinux.org")
	fmt.Println("# https://archlinuxarm.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Arch_Linux")
	fmt.Println("# https://en.wikipedia.org/wiki/Arch_Linux_ARM")
	fmt.Println("# https://distrowatch.com/arch")

	dumpOne("https://mirror.xenyth.net/archlinux/iso/latest")

	// Get the signing keys too
	fmt.Println("https://keys.openpgp.org/vks/v1/by-fingerprint/3E80CA1A8B89F69CBA57D98A76A5EF9054449A5C")
	fmt.Println("	dir=Linux/Arch")
	fmt.Println("	out=3E80CA1A8B89F69CBA57D98A76A5EF9054449A5C.pgp.asc")
	fmt.Println("https://keys.openpgp.org/vks/v1/by-fingerprint/68B3537F39A313B3E574D06777193F152BDBE6A6")
	fmt.Println("	dir=Linux/Arch")
	fmt.Println("	out=68B3537F39A313B3E574D06777193F152BDBE6A6.pgp.asc")
}

func dumpOne(url string) {
	fmt.Println(fmt.Sprintf("# %s", url))

	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Do a first pass to get the version number to use when renaming the checksum files
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := ""
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "archlinux") && !strings.Contains(href, "archlinux-x86_64") && !strings.Contains(href, "archlinux-bootstrap") && !strings.Contains(href, ".torrent") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Linux/Arch")
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
				}
			} else if !strings.Contains(href, "sums.txt") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})

	// Now that we know the release number, we can give the generic checksum files sensible names
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "sums.txt") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Linux/Arch")
				fmt.Println(fmt.Sprintf("	out=archlinux-%s-x86_64-%s", ver, href))
			}
		}
	})
}
