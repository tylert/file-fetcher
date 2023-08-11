package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
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

func doIt() {
	res, err := http.Get("https://salsa.debian.org/api/v4/projects/69786/releases/permalink/latest")
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

	fmt.Println("# https://salsa.debian.org/jgoerzen/filespooler")
	fmt.Println("# https://www.complete.org/filespooler")

	// This project uses version strings that start with "v" in some places
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := reg.FindString(rel.TagName)

	// Compiled binaries
	for i := 0; i < len(rel.Assets.Links); i++ {
		fmt.Println(fmt.Sprintf("# %s", rel.Assets.Links[i].Name))
		fmt.Println(rel.Assets.Links[i].DirectAssetURL)
		fmt.Println("	dir=Filespooler")
		fmt.Println(fmt.Sprintf("	out=fspl-%s", ver))
	}

	// Source code
	for j := 0; j < len(rel.Assets.Sources); j++ {
		fmt.Println(rel.Assets.Sources[j].URL)
		fmt.Println("	dir=Filespooler")
	}
}

func main() {
	doIt()
}
