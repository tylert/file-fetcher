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
	res, err := http.Get("https://api.github.com/repos/siderolabs/talos/releases/latest")
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
	fmt.Println("# https://github.com/siderolabs/talos/releases")
	fmt.Println("# https://github.com/siderolabs/talos")
	fmt.Println("# https://talos.dev")
	fmt.Println("# https://siderolabs.com")
	fmt.Println("# https://siderolabs.com/platform/talos-os-for-kubernetes")
	fmt.Println("# https://kubito.dev/posts/talos-linux-raspberry-pi")

	// Compiled binaries
	for i := 0; i < len(rel.Assets); i++ {
		if strings.Contains(rel.Assets[i].Name, "metal-amd64.iso") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Talos")
			fmt.Println(fmt.Sprintf("	out=talos-%s-metal-amd64.iso", ver))
		} else if strings.Contains(rel.Assets[i].Name, "metal-amd64.raw.xz") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Talos")
			fmt.Println(fmt.Sprintf("	out=talos-%s-metal-amd64.raw.xz", ver))
		} else if strings.Contains(rel.Assets[i].Name, "talosctl-linux-amd64") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Talos")
			fmt.Println(fmt.Sprintf("	out=talosctl-%s-linux-amd64", ver))
		} else if strings.Contains(rel.Assets[i].Name, "talosctl-linux-arm64") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Talos")
			fmt.Println(fmt.Sprintf("	out=talosctl-%s-linux-arm64", ver))
		} else if strings.Contains(rel.Assets[i].Name, "sha512sum.txt") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Talos")
			fmt.Println(fmt.Sprintf("	out=talos-%s-sha512sum.txt", ver))
		} else {
			fmt.Println(fmt.Sprintf("# skipped %s", rel.Assets[i].Name))
		}
	}

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=Talos")
	fmt.Println(fmt.Sprintf("	out=talos-%s-src.tar.gz", ver))
}

func main() {
	doIt()
}
