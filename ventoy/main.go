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
		log.Fatalf("Kaboom!")
	}

	fmt.Println("# https://github.com/ventoy/Ventoy")
	fmt.Println("# https://www.ventoy.net")

	// Source code
	fmt.Println(fmt.Sprintf("%s", rel.TarballURL))
	fmt.Println("	auto-file-renaming=false")
	fmt.Println("	conditional-get=true")
	fmt.Println("	dir=Ventoy")
	fmt.Println(fmt.Sprintf("	out=ventoy-%s-src.tar.gz", rel.TagName))

	// Compiled binaries
	for i := 0; i < len(rel.Assets); i++ {
		if strings.Contains(rel.Assets[i].Name, "-linux") {
			fmt.Println(fmt.Sprintf("%s", rel.Assets[i].BrowserDownloadURL))
			fmt.Println("	auto-file-renaming=false")
			fmt.Println("	conditional-get=true")
			fmt.Println("	dir=Ventoy")
		}
		if strings.Contains(rel.Assets[i].Name, "sha256.txt") {
			fmt.Println(fmt.Sprintf("%s", rel.Assets[i].BrowserDownloadURL))
			fmt.Println("	auto-file-renaming=false")
			fmt.Println("	dir=Ventoy")
			fmt.Println(fmt.Sprintf("	out=ventoy-%s-sha256.txt", rel.TagName))
		}
	}
}
