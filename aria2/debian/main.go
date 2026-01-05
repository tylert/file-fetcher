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
	Debian()
}

func Debian() {
	// Spit out some handy links
	fmt.Println("# https://cdimage.debian.org/cdimage")
	fmt.Println("# https://debian.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Debian_version_history#Release_table")
	fmt.Println("# https://en.wikipedia.org/wiki/Debian")
	fmt.Println("# https://distrowatch.com/debian")

	// Live CD versions
	dumpOne("https://cdimage.debian.org/cdimage/release/current-live/amd64/iso-hybrid", "xfce")

	// Spit out some more handy links
	// fmt.Println("# https://ftp-master.debian.org/keys.html")
	// fmt.Println("# https://ftp-master.debian.org/keys")

	// Get the signing keys too
	// fmt.Println("https://ftp-master.debian.org/keys/release-13.asc") // 41587F7DB8C774BCCF131416762F67A0B2C39DE4
	// fmt.Println("	dir=Linux/Debian")
}

func dumpOne(url string, variant string) {
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
	if res.StatusCode != http.StatusOK { // 200
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Do a first pass to get the version number to use when renaming the checksum files
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := ""
	doc.Find("td.indexcolname a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if variant != "" {
				if strings.Contains(href, variant) && !strings.Contains(href, "contents") && !strings.Contains(href, "log") && !strings.Contains(href, "packages") {
					fmt.Println(fmt.Sprintf("%s/%s", url, href))
					fmt.Println("	dir=Linux/Debian")
					if reg.FindString(href) != "" {
						ver = reg.FindString(href)
					}
				}
			} else {
				if strings.Contains(href, "netinst.iso") && !strings.Contains(href, "-edu-") && !strings.Contains(href, "-mac-") {
					fmt.Println(fmt.Sprintf("%s/%s", url, href))
					fmt.Println("	dir=Linux/Debian")
					if reg.FindString(href) != "" {
						ver = reg.FindString(href)
					}
				}
			}
		}
	})

	// If there's no version string, that means it's a pending release (i.e.:  "testing")
	if ver == "" {
		ver = "testing"
	}

	// Now that we know the release number, we can give the generic checksum files sensible names
	doc.Find("td.indexcolname a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "SHA") && !strings.Contains(href, "SHA1SUMS") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Linux/Debian")
				fmt.Println(fmt.Sprintf("	out=debian-%s-%s", ver, href))
			}
		}
	})
}
