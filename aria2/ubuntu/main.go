package main

import (
	"fmt"
	"log"
	"net/http"
	//"regexp"
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

	// XXX FIXME TODO  Fix the names of the checksum files!!!
	// Do a first pass to get the version number to use when renaming the checksum files
	//reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	//ver := ""
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if (strings.Contains(href, ".iso") || strings.Contains(href, ".zsync") || strings.Contains(href, ".list") || strings.Contains(href, ".manifest") || strings.Contains(href, "SHA256SUMS")) && !strings.Contains(href, ".torrent") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Ubuntu")
			}
		}
	})
}

func main() {
	fmt.Println("# https://mirror.xenyth.net/ubuntu-releases")
	fmt.Println("# https://releases.ubuntu.com")
	fmt.Println("# https://cdimage.ubuntu.com")
	fmt.Println("# https://en.wikipedia.org/wiki/Ubuntu")
	fmt.Println("# https://en.wikipedia.org/wiki/Ubuntu_version_history")
	fmt.Println("# https://distrowatch.com/ubuntu")

	fmt.Println("# https://ubuntu.com/download/raspberry-pi")

	// XXX FIXME TODO  Try to determine the list of "current" releases programatically!!!
	dumpOne("https://mirror.xenyth.net/ubuntu-releases/23.04") // lunar
	dumpOne("https://mirror.xenyth.net/ubuntu-releases/22.04") // jammy
	dumpOne("https://mirror.xenyth.net/ubuntu-releases/20.04") // focal
}
