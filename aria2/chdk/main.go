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
	fmt.Println("# https://forum.chdk-treff.de/download.php")     // stabile
	fmt.Println("# https://forum.chdk-treff.de/download_dev.php") // Vorshauversion
	fmt.Println("# https://forum.chdk-treff.de/index.php")        // de forum
	fmt.Println("# https://chdk.setepontos.com")                  // en forum
	fmt.Println("# https://chdk.fandom.com/wiki/CHDK")

	// XXX FIXME TODO  Workaround the dumb new clicky download thing;  Just give me the links!!!
	//                 "Digital IXUS Serie" -> "IXUS132 (ELPH115)" -> "100b"
	// dumpOne("https://mighty-hoernsche.de", "_elph115-")       // stable
	// dumpOne("https://mighty-hoernsche.de/trunk", "_elph115-") // unstable
}

func main() {
	doIt()
}
