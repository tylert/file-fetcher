package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Release struct {
	HTMLURL         string    `json:"html_url"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []struct {
		Name               string    `json:"name"`
		ContentType        string    `json:"content_type"`
		Size               int       `json:"size"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadURL string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballURL string `json:"tarball_url"`
	ZipballURL string `json:"zipball_url"`
}

func dumpBin() {
	res, err := http.Get("https://libreelec.tv/downloads/raspberry")
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

	// Compiled binaries
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, ".gz") {
				fmt.Println(href)
				fmt.Println("	dir=LibreELEC")
				return false
			}
		}
		return true
	})
}

func dumpSrc() {
	res, err := http.Get("https://api.github.com/repos/LibreELEC/LibreELEC.tv/releases/latest")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	var rel Release
	err = json.NewDecoder(res.Body).Decode(&rel)
	if err != nil {
		log.Fatalf("Kaboom!")
	}

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=LibreELEC")
	// XXX FIXME TODO  Fix the filename for the source tarballs!!!
}

func main() {
	fmt.Println("# https://github.com/LibreELEC/LibreELEC.tv")
	fmt.Println("# https://libreelec.tv/downloads/raspberry")
	fmt.Println("# https://libreelec.tv")
	fmt.Println("# https://en.wikipedia.org/wiki/LibreELEC")

	dumpBin()
	dumpSrc()
}
