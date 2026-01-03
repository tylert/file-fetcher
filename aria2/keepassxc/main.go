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
	res, err := client.Get("https://api.github.com/repos/keepassxreboot/keepassxc/releases/latest")
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
	fmt.Println("# https://github.com/keepassxreboot/keepassxc")
	fmt.Println("# https://github.com/keepassxreboot/keepassxc/releases")
	fmt.Println("# https://keepassxc.org/download")
	fmt.Println("# https://keepassxc.org")
	fmt.Println("# https://en.wikipedia.org/wiki/KeePassXC")

	// Compiled binaries and source code
	for i := 0; i < len(rel.Assets); i++ {
		if strings.Contains(rel.Assets[i].Name, ".AppImage") || strings.Contains(rel.Assets[i].Name, "-src") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=KeePassXC")
		} else {
			fmt.Println(fmt.Sprintf("# skipped %s", rel.Assets[i].Name))
		}
	}
	// XXX FIXME TODO  There is rather poor consistency in the filename case in this project
}
