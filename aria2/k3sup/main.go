/*usr/bin/env go run "$0" "$@"; exit;*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://api.github.com/repos/alexellis/k3sup/releases/latest")
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

	// Spit out some handy links
	fmt.Println("# https://github.com/alexellis/k3sup")
	fmt.Println("# https://github.com/alexellis/k3sup/releases")

	// Compiled binaries
	for i := 0; i < len(rel.Assets); i++ {
		if !strings.Contains(rel.Assets[i].Name, "-armhf") && !strings.Contains(rel.Assets[i].Name, "-darwin") && !strings.Contains(rel.Assets[i].Name, ".exe") {
			if strings.Contains(rel.Assets[i].Name, "-arm64.sha256") {
				fmt.Println(rel.Assets[i].BrowserDownloadURL)
				fmt.Println("	dir=k3sup")
				fmt.Println(fmt.Sprintf("	out=k3sup-%s-linux-arm64.sha256.txt", rel.TagName))
			} else if strings.Contains(rel.Assets[i].Name, "-arm64") {
				fmt.Println(rel.Assets[i].BrowserDownloadURL)
				fmt.Println("	dir=k3sup")
				fmt.Println(fmt.Sprintf("	out=k3sup-%s-linux-arm64", rel.TagName))
			} else if strings.Contains(rel.Assets[i].Name, ".sha256") {
				fmt.Println(rel.Assets[i].BrowserDownloadURL)
				fmt.Println("	dir=k3sup")
				fmt.Println(fmt.Sprintf("	out=k3sup-%s-linux-amd64.sha256.txt", rel.TagName))
			} else {
				fmt.Println(rel.Assets[i].BrowserDownloadURL)
				fmt.Println("	dir=k3sup")
				fmt.Println(fmt.Sprintf("	out=k3sup-%s-linux-amd64", rel.TagName))
			}
		} else {
			fmt.Println(fmt.Sprintf("# skipped %s", rel.Assets[i].Name))
		}
	}

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=k3sup")
	fmt.Println(fmt.Sprintf("	out=k3sup-%s-src.tar.gz", rel.TagName))
}
