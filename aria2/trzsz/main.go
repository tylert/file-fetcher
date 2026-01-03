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

func dumpOne(url string, tool string) {
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

	var rel Release
	err = json.NewDecoder(res.Body).Decode(&rel)
	if err != nil {
		log.Fatal(err)
	}

	// This project uses version strings that start with "v" in some places
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := reg.FindString(rel.TagName)

	// Compiled binaries
	for i := 0; i < len(rel.Assets); i++ {
		if (strings.Contains(rel.Assets[i].Name, "linux") || strings.Contains(rel.Assets[i].Name, "checksums")) && !strings.Contains(rel.Assets[i].Name, "386") && !strings.Contains(rel.Assets[i].Name, "rpm") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=trzsz")
		} else {
			fmt.Println(fmt.Sprintf("# skipped %s", rel.Assets[i].Name))
		}
	}

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=trzsz")
	fmt.Println(fmt.Sprintf("	out=%s_%s_src.tar.gz", tool, ver))
}

func main() {
	// Spit out some handy links
	fmt.Println("# https://github.com/trzsz/trzsz-go")
	fmt.Println("# https://github.com/trzsz/trzsz-go/releases")
	fmt.Println("# https://github.com/trzsz/trzsz-ssh")
	fmt.Println("# https://github.com/trzsz/trzsz-ssh/releases")
	fmt.Println("# https://github.com/trzsz/trzsz.github.io")
	fmt.Println("# https://trzsz.github.io/go")
	fmt.Println("# https://trzsz.github.io/ssh")
	fmt.Println("# https://trzsz.github.io")

	dumpOne("https://api.github.com/repos/trzsz/trzsz-go/releases/latest", "trzsz")
	dumpOne("https://api.github.com/repos/trzsz/trzsz-ssh/releases/latest", "tssh")

	// XXX FIXME TODO  Fetch the https://github.com/trzsz/trzsz.github.io docs!!!
}
