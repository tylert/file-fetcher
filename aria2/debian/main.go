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
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := ""
	doc.Find("td.indexcolname a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "netinst.iso") && !strings.Contains(href, "-edu-") && !strings.Contains(href, "-mac-") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Debian")
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
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
				fmt.Println("	dir=Debian")
				fmt.Println(fmt.Sprintf("	out=debian-%s-%s", ver, href))
			}
		}
	})
}

func doIt() {
	// Spit out some handy links
	fmt.Println("# https://cdimage.debian.org/cdimage")
	fmt.Println("# https://debian.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Debian_version_history#Release_table")
	fmt.Println("# https://en.wikipedia.org/wiki/Debian")
	fmt.Println("# https://distrowatch.com/debian")

	dumpOne("https://cdimage.debian.org/cdimage/weekly-builds/amd64/iso-cd")               // testing
	dumpOne("https://cdimage.debian.org/cdimage/release/current/amd64/iso-cd")             // stable
	dumpOne("https://cdimage.debian.org/cdimage/archive/latest-oldstable/amd64/iso-cd")    // oldstable
	dumpOne("https://cdimage.debian.org/cdimage/archive/latest-oldoldstable/amd64/iso-cd") // oldoldstable

	// Spit out some more handy links
	fmt.Println("# https://ftp-master.debian.org/keys.html")
	fmt.Println("# https://ftp-master.debian.org/keys")

	// Get the signing keys too
	fmt.Println("https://ftp-master.debian.org/keys/release-12.asc") // 4D64FEC119C2029067D6E791F8D2585B8783D481
	fmt.Println("	dir=Debian")
	fmt.Println("https://ftp-master.debian.org/keys/release-11.asc") // A4285295FC7B1A81600062A9605C66F00D6C9793
	fmt.Println("	dir=Debian")
	fmt.Println("https://ftp-master.debian.org/keys/release-10.asc") // 6D33866EDD8FFA41C0143AEDDCC9EFBF77E11517
	fmt.Println("	dir=Debian")
}

func main() {
	doIt()
}
