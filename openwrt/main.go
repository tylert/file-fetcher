package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	// "strings"
	// "github.com/PuerkitoBio/goquery"
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
	TarballURL      string    `json:"tarball_url"`
}

// curl https://api.github.com/repos/openwrt/openwrt/releases/latest | jq '.tarball_url'
// curl https://api.github.com/repos/openwrt/openwrt/releases/latest | jq '.name'
// https://downloads.openwrt.org/releases/${VERSION}/targets/ath79/generic/openwrt-${VERSION}-ath79-generic-tplink_eap225-v3-{initramfs-kernel,squashrs-factory,squashfs-sysupgrade}.bin
// https://downloads.openwrt.org/releases/${VERSION}/targets/ramips/mt7621/openwrt-${VERSION}-ramips-mt7621-ubnt_edgerouter-x-{initramfs-kernel,squashfs-sysupgrade}.bin
// + sha256sums

func main() {
	res, err := http.Get("https://api.github.com/repos/openwrt/openwrt/releases/latest")
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

	fmt.Println(fmt.Sprintf("%s", rel.TarballURL))
	fmt.Println(fmt.Sprintf("https://downloads.openwrt.org/releases/%s/targets/ath79/generic/", rel.Name))
	fmt.Println(fmt.Sprintf("https://downloads.openwrt.org/releases/%s/targets/ramips/mt7621/", rel.Name))
}
