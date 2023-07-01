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

func dumpBin(url string, target string) {
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
		log.Fatalf("Error loading HTTP response body.", err)
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		moo := s.Find("td.n a")
		poo := s.Find("td.sh").Text()
		href, ok := moo.Attr("href")
		if ok {
			if strings.Contains(href, target) && !strings.Contains(href, "-sfp-") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=OpenWRT")
				fmt.Println(fmt.Sprintf("	checksum=sha-256=%s", poo))
			}
		}
	})
}

func main() {
	res, err := http.Get("https://api.github.com/repos/openwrt/openwrt/releases/latest")
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

	fmt.Println("# https://github.com/openwrt/openwrt")
	fmt.Println("# https://openwrt.org")
	fmt.Println("# https://downloads.openwrt.org/releases")
	fmt.Println("# https://firmware-selector.openwrt.org")
	fmt.Println("# https://en.wikipedia.org/wiki/OpenWrt")

	// Compiled binaries
	dumpBin(fmt.Sprintf("https://downloads.openwrt.org/releases/%s/targets/ath79/generic", rel.Name), "tplink_eap245-v3")
	dumpBin(fmt.Sprintf("https://downloads.openwrt.org/releases/%s/targets/ramips/mt7621", rel.Name), "ubnt_edgerouter-x")

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=OpenWRT")
	fmt.Println(fmt.Sprintf("	out=openwrt-%s-src.tar.gz", rel.Name))
}
