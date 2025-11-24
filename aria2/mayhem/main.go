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
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://api.github.com/repos/portapack-mayhem/mayhem-firmware/releases/latest")
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
	fmt.Println("# https://github.com/portapack-mayhem/mayhem-firmware")
	fmt.Println("# https://github.com/portapack-mayhem/mayhem-firmware/releases")
	fmt.Println("# https://rtl-sdr.com/a-review-of-the-new-hackrf-portapack-h4m")

	// Compiled binaries
	for i := 0; i < len(rel.Assets); i++ {
		if strings.Contains(rel.Assets[i].Name, "FIRMWARE.zip") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Mayhem")
			thingy := strings.ReplaceAll(rel.Assets[i].Name, fmt.Sprintf("v%s", ver), ver)
			fmt.Println(fmt.Sprintf("	out=%s", thingy))
		} else if strings.Contains(rel.Assets[i].Name, "COPY_TO_SDCARD.zip") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Mayhem")
			thingy := strings.ReplaceAll(rel.Assets[i].Name, fmt.Sprintf("v%s", ver), ver)
			fmt.Println(fmt.Sprintf("	out=%s", thingy))
		} else if strings.Contains(rel.Assets[i].Name, "ppfw.tar") {
			fmt.Println(rel.Assets[i].BrowserDownloadURL)
			fmt.Println("	dir=Mayhem")
			thingy := strings.ReplaceAll(rel.Assets[i].Name, fmt.Sprintf("v%s", ver), ver)
			fmt.Println(fmt.Sprintf("	out=%s", thingy))
		} else {
			fmt.Println(fmt.Sprintf("# skipped %s", rel.Assets[i].Name))
		}
	}

	// Source code
	fmt.Println(rel.TarballURL)
	fmt.Println("	dir=Mayhem")
	fmt.Println(fmt.Sprintf("	out=mayhem_%s_SOURCE.tar.gz", ver))
}

func main() {
	doIt()
}
