package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func dumpOne(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body.", err)
	}

	fmt.Println(fmt.Sprintf("# %s", url))

	// Do a first pass to get the version number to use when renaming the checksum files
	reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	ver := ""
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "archlinux") && !strings.Contains(href, "archlinux-x86_64") && !strings.Contains(href, "archlinux-bootstrap-x86_64") && !strings.Contains(href, ".torrent") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=ArchLinux")
				if reg.FindString(href) != "" {
					ver = reg.FindString(href)
				}
			} else if !strings.Contains(href, "sums.txt") {
				fmt.Println(fmt.Sprintf("# skipped %s", href))
			}
		}
	})
	// Now that we know the release number, we can give the generic checksum files sensible names
	doc.Find("div.name a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			if strings.Contains(href, "sums.txt") {
				fmt.Println(fmt.Sprintf("%s/%s", url, href))
				fmt.Println("	dir=ArchLinux")
				fmt.Println(fmt.Sprintf("	out=archlinux-%s-x86_64-%s", ver, href))
			}
		}
	})
}

func main() {
	fmt.Println("# https://mirror.xenyth.net/archlinux")
	fmt.Println("# https://archlinux.org/mirrors")
	fmt.Println("# https://archlinux.org/download")
	fmt.Println("# https://archlinuxarm.org/about/downloads")
	fmt.Println("# https://archlinux.org")
	fmt.Println("# https://archlinuxarm.org")
	fmt.Println("# https://archlinuxarm.org/platforms/armv8/broadcom/raspberry-pi-4")
	fmt.Println("# https://en.wikipedia.org/wiki/Arch_Linux")
	fmt.Println("# https://en.wikipedia.org/wiki/Arch_Linux_ARM")
	fmt.Println("# https://distrowatch.com/arch")

	dumpOne("https://mirror.xenyth.net/archlinux/iso/latest")

	// Get the Pi4 image too
	fmt.Println("http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-aarch64-latest.tar.gz")
	fmt.Println("	dir=ArchLinux")
	//fmt.Println(fmt.Sprintf("	out=ArchLinuxARM-rpi-aarch64-%s.tar.gz", today))
	fmt.Println("http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-aarch64-latest.tar.gz.md5")
	fmt.Println("	dir=ArchLinux")
	//fmt.Println(fmt.Sprintf("	out=ArchLinuxARM-rpi-aarch64-%s.tar.gz.md5", today))
	fmt.Println("http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-aarch64-latest.tar.gz.sig")
	fmt.Println("	dir=ArchLinux")
	//fmt.Println(fmt.Sprintf("	out=ArchLinuxARM-rpi-aarch64-%s.tar.gz.sig", today))
}