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
	ArtixLinux()
}

func ArtixLinux() {
	// Spit out some handy links
	fmt.Println("# https://iso.artixlinux.org/isos.php")
	fmt.Println("# https://artixlinux.org")
	fmt.Println("# https://en.wikipedia.org/wiki/Artix_Linux")
	fmt.Println("# https://distrowatch.com/artix")

	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://iso.artixlinux.org/isos.php")
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
	reg := regexp.MustCompile(`\d+\d+\d+\d+\d+\d+\d+\d+`)
	ver := ""
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "artix-xfce-s6") {
				fmt.Println(fmt.Sprintf("%s", href))
				fmt.Println("	dir=Linux/Artix")
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
				}
			}
		}
	})

	// Now that we know the release number, we can give the generic checksum files sensible names
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "sha256sums") {
				fmt.Println(fmt.Sprintf("%s", href))
				fmt.Println("	dir=Linux/Artix")
				fmt.Println(fmt.Sprintf("	out=artix-sha256sums-%s-x86_64.txt", ver))
			}
		}
	})
}
