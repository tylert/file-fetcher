package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func doIt() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://grapheneos.org/releases")
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

	// Spit out some handy links
	fmt.Println("# https://grapheneos.org")
	fmt.Println("# https://grapheneos.org/releases")
	fmt.Println("# https://comparisontabl.es/google-pixel-phones")
	fmt.Println("# https://en.wikipedia.org/wiki/GrapheneOS")

	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "https")
	}

	count := 3
	doc.Find("a").FilterFunction(f).EachWithBreak(func(_ int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")

		if strings.Contains(link, "caiman") {
			fmt.Printf("%s\n", link)
			fmt.Println("	dir=Android/Google_Pixel_9_Pro_GEC77_caiman/GrapheneOS")
			count--
			if count <= 0 {
				return false
			}
		}
		return true
	})
}

func main() {
	doIt()
}
