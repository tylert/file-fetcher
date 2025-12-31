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
)

type Release struct {
	Assets struct {
		Links []struct {
			DirectAssetURL string `json:"direct_asset_url"`
			Name           string `json:"name"`
			URL            string `json:"url"`
		} `json:"links"`
		Sources []struct {
			URL string `json:"url"`
		} `json:"sources"`
	} `json:"assets"`
	CreatedAt  time.Time `json:"created_at"`
	ReleasedAt time.Time `json:"released_at"`
	TagName    string    `json:"tag_name"`
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://salsa.debian.org/api/v4/projects/69786/releases/permalink/latest")
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

	// This project uses version strings that start with "v" in some places
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := reg.FindString(rel.TagName)

	// Spit out some handy links
	fmt.Println("# https://salsa.debian.org/jgoerzen/filespooler/-/releases")
	fmt.Println("# https://salsa.debian.org/jgoerzen/filespooler")
	fmt.Println("# https://complete.org/filespooler")

	// Compiled binaries
	for i := 0; i < len(rel.Assets.Links); i++ {
		if strings.Contains(rel.Assets.Links[i].Name, "Linux x86_64") {
			fmt.Println(fmt.Sprintf("# %s", rel.Assets.Links[i].Name))
			fmt.Println(rel.Assets.Links[i].DirectAssetURL)
			fmt.Println("	dir=Filespooler")
			fmt.Println(fmt.Sprintf("	out=fspl-%s-linux-amd64", ver))
		} else if strings.Contains(rel.Assets.Links[i].Name, "Linux aarch64") {
			fmt.Println(fmt.Sprintf("# %s", rel.Assets.Links[i].Name))
			fmt.Println(rel.Assets.Links[i].DirectAssetURL)
			fmt.Println("	dir=Filespooler")
			fmt.Println(fmt.Sprintf("	out=fspl-%s-linux-arm64", ver))
		} else {
			whatzit := strings.Split(rel.Assets.Links[i].DirectAssetURL, "/")
			fmt.Println(fmt.Sprintf("# skipped %s", whatzit[len(whatzit)-3]))
		}
	}

	// Source code
	for j := 0; j < len(rel.Assets.Sources); j++ {
		if strings.Contains(rel.Assets.Sources[j].URL, ".tar.gz") {
			fmt.Println(rel.Assets.Sources[j].URL)
			fmt.Println("	dir=Filespooler")
			fmt.Println(fmt.Sprintf("	out=filespooler-%s-src.tar.gz", ver))
		} else {
			whatzit := strings.Split(rel.Assets.Sources[j].URL, "/")
			fmt.Println(fmt.Sprintf("# skipped %s", whatzit[len(whatzit)-1]))
		}
	}
}
