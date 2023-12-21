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

func doIt() {
	res, err := http.Get("https://api.github.com/repos/ventoy/Ventoy/releases/latest")
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
	fmt.Println("# https://github.com/ventoy/Ventoy/releases")
	fmt.Println("# https://github.com/ventoy/Ventoy")
	fmt.Println("# https://ventoy.net/en/download.html")
	fmt.Println("# https://ventoy.net")
	fmt.Println("# https://en.wikipedia.org/wiki/Ventoy")

	// Compiled binaries
	for i := 0; i < len(rel.Assets); i++ {
		if strings.Contains(rel.Assets[i].Name, "-linux") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Ventoy")
		} else if strings.Contains(rel.Assets[i].Name, "sha256.txt") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Ventoy")
			fmt.Println(fmt.Sprintf("	out=ventoy-%s-sha256.txt", ver))
		} else {
			fmt.Println(fmt.Sprintf("# skipped %s", rel.Assets[i].Name))
		}
	}

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=Ventoy")
	fmt.Println(fmt.Sprintf("	out=ventoy-%s-src.tar.gz", ver))
}

func main() {
	doIt()
}
