package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
  "encoding/json"
	// "github.com/PuerkitoBio/goquery"
)

type release struct {
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
