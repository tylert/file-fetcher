/*usr/bin/env go run "$0" "$@"; exit;*/
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

type Entry []struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	NodeID string `json:"node_id"`
}

func dumpSrc(url string, target string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	var tag Entry
	err = json.NewDecoder(res.Body).Decode(&tag)
	if err != nil {
		log.Fatal(err)
	}

	// This project uses version strings that start with "v" in some places
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := reg.FindString(tag[0].Name)

	// Assume that the first hit is the newest and then stop after that
	fmt.Println(fmt.Sprintf("%s", tag[0].TarballURL))
	fmt.Println("	dir=Editcp")
	fmt.Println(fmt.Sprintf("	out=%s-%s-src.tar.gz", target, ver))
}

func dumpBin(url string, target string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(url)
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

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, target) {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=Editcp")
			}
		}
	})
}

func doIt() {
	// Spit out some handy links
	fmt.Println("# https://farnsworth.org/dale/codeplug/dmrRadio/downloads")
	fmt.Println("# https://farnsworth.org/dale/codeplug/editcp/downloads")
	fmt.Println("# https://farnsworth.org/dale/codeplug/editcp")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/codeplug")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/debug")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/dfu")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/dmrRadio")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/docCodeplug")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/docker")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/docs")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/editcp")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/genCodeplugInfo")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/genFileData")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/stdfu")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/ui")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/userdb")

	// Compiled binaries
	dumpBin("https://farnsworth.org/dale/codeplug/dmrRadio/downloads/linux", "dmrRadio")
	dumpBin("https://farnsworth.org/dale/codeplug/editcp/downloads/linux", "editcp")

	// Source code
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/codeplug/tags", "codeplug")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/debug/tags", "debug")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/dfu/tags", "dfu")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/dmrRadio/tags", "dmrRadio")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/docCodeplug/tags", "docCodeplug")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/docker/tags", "docker")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/docs/tags", "docs")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/editcp/tags", "editcp")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/genCodeplugInfo/tags", "genCodeplugInfo")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/genFileData/tags", "genFileData")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/stdfu/tags", "stdfu")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/ui/tags", "ui")
	dumpSrc("https://api.github.com/repos/dalefarnsworth-dmr/userdb/tags", "userdb")
}

func main() {
	doIt()
}
