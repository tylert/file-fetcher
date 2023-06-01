package main

import (
	"fmt"
	"log"
	"net/http"
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

	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if !strings.Contains(href, ".torrent") && !strings.Contains(href, "MD5") && !strings.Contains(href, "SHA1SUM") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	allow-overwrite=true")
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	dir=CentOS")
				fmt.Println("	file-allocation=falloc")
			}
		}
	})
}

func main() {
	dumpOne("https://mirror.xenyth.net/centos-stream/9-stream/BaseOS/x86_64/iso")
	dumpOne("https://mirror.xenyth.net/centos/8-stream/isos/x86_64")
	dumpOne("https://mirror.xenyth.net/centos/7/isos/x86_64")
}
