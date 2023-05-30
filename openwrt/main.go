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
		log.Fatalf("Error loading HTTP response body.", err)
	}

	// XXX FIXME TODO  Extract the corresponding sha256sum along with the filename!!!
	// <tr><td class="n">
	// <a href="openwrt-22.03.5-bla-bla-bla.bin">bla-bla-bla.bin</a>
	// </td>
	// <td class="sh">3b28c9bf308b38ccb95aadbd4c52d9c686b8af6ba9ad1b00694f7fedd1f7506f</td>

	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if strings.Contains(href, "tplink_eap225-v3") || (strings.Contains(href, "ubnt_edgerouter-x") && !strings.Contains(href, "-sfp-")) {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	auto-file-renaming=false")
				fmt.Println("	dir=OpenWRT")
				fmt.Println("	file-allocation=falloc")
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

	// Compiled binaries
	dumpOne(fmt.Sprintf("https://downloads.openwrt.org/releases/%s/targets/ath79/generic", rel.Name))
	dumpOne(fmt.Sprintf("https://downloads.openwrt.org/releases/%s/targets/ramips/mt7621", rel.Name))

	// Source code
	fmt.Println(fmt.Sprintf("%s", rel.TarballURL))
	fmt.Println("	auto-file-renaming=false")
	fmt.Println("	dir=OpenWRT")
	fmt.Println("	file-allocation=falloc")
	fmt.Println(fmt.Sprintf("	out=openwrt-%s-src.tar.gz", rel.Name))
}
