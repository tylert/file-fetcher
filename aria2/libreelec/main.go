package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
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

func dumpBin() string {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://libreelec.tv/downloads/raspberry")
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

	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := ""

	// Do a first pass to get the version number to use for the source code
	// The newest version number will be the first one we find (top-down)
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, ".gz") {
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
				}
				return false
			}
		}
		return true
	})

	// Compiled binaries (stop after finding exactly 1)
	count := 1
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, ver) && strings.Contains(href, ".gz") && strings.Contains(href, "RPi4") {
				fmt.Println(href)
				fmt.Println("	dir=Linux/LibreELEC")
				count--
				if count <= 0 {
					return false
				}
			} else if strings.Contains(href, ".gz") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
		return true
	})

	return ver
}

func dumpSrc(ver string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://api.github.com/repos/LibreELEC/LibreELEC.tv/releases/latest")
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
		log.Fatal(err)
	}

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=Linux/LibreELEC")
	fmt.Println(fmt.Sprintf("	out=LibreELEC-%s-src.tar.gz", ver))
}

func doIt() {
	// Spit out some handy links
	fmt.Println("# https://github.com/LibreELEC/LibreELEC.tv")
	fmt.Println("# https://github.com/LibreELEC/LibreELEC.tv/releases")
	fmt.Println("# https://libreelec.tv")
	fmt.Println("# https://libreelec.tv/downloads/raspberry")
	fmt.Println("# https://en.wikipedia.org/wiki/LibreELEC")

	ver := dumpBin()
	dumpSrc(ver)
}

func main() {
	doIt()
}
