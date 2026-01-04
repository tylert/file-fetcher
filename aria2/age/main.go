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

func main() {
	Age()
}

func Age() {
	// Spit out some handy links
	fmt.Println("# https://github.com/FiloSottile/age")
	fmt.Println("# https://github.com/FiloSottile/age/releases")
	fmt.Println("# https://github.com/FiloSottile/awesome-age")
	fmt.Println("# https://age-encryption.org")
	fmt.Println("# https://complete.org/age-encryption")
	fmt.Println("# https://words.filippo.io/dispatches/age-authentication")
	fmt.Println("# https://yaeba.github.io/blog/age")

	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://api.github.com/repos/FiloSottile/age/releases/latest")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	var rel Release
	if err = json.NewDecoder(res.Body).Decode(&rel); err != nil {
		log.Fatal(err)
	}

	// This project uses version strings that start with "v" in some places
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := reg.FindString(rel.TagName)

	// Compiled binaries
	for i := 0; i < len(rel.Assets); i++ {
		if strings.Contains(rel.Assets[i].Name, "-linux") && strings.Contains(rel.Assets[i].Name, "64") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=age")
			thingy := strings.ReplaceAll(rel.Assets[i].Name, fmt.Sprintf("v%s", ver), ver)
			fmt.Println(fmt.Sprintf("	out=%s", thingy))
		} else {
			fmt.Println(fmt.Sprintf("# skipped %s", rel.Assets[i].Name))
		}
	}

	// Documentation
	fmt.Println("https://raw.githubusercontent.com/FiloSottile/age/main/doc/age-keygen.1.html")
	fmt.Println("	dir=age")
	fmt.Println("https://raw.githubusercontent.com/FiloSottile/age/main/doc/age.1.html")
	fmt.Println("	dir=age")

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=age")
	fmt.Println(fmt.Sprintf("	out=age-%s-src.tar.gz", ver))
}
