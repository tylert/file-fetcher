package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func dumpOne(url string, target string) {
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

	// Break after we find a single match (e.g.:  device twins)
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			// Don't bother with the partial builds, only full ones
			if strings.Contains(href, target) && strings.Contains(href, "full.zip") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=CHDK")
				return false
			}
		}
		return true
	})
}

func doIt() {
	// Spit out some handy links
	fmt.Println("# https://en.wikipedia.org/wiki/DIGIC#CHDK")
	fmt.Println("# https://chdk.fandom.com/wiki/Downloads")
	fmt.Println("# https://build.chdk.photos")
	fmt.Println("# https://app.assembla.com/spaces/chdk/subversion/source")

	// Go to https://build.chdk.photos and click on the following:
	//     Development / Unstable (trunk)
	//     "Release / Stable (release)
	//         -> Digital IXUS (SD, ELPH, IXY)
	//             -> IXUS 132 (ELPH 115 IS, IXY 90F)
	//                 -> 100b
	//                     -> Complete build (zip link with inline sha256 checksum)
}

func main() {
	doIt()
}
