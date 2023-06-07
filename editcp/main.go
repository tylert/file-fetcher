package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"regexp"
)

type Entry []struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	NodeID string `json:"node_id"`
}

func dumpOne(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	var tag Entry
	err = json.NewDecoder(res.Body).Decode(&tag)
	if err != nil {
		log.Fatalf("Kaboom!")
	}

	// This project uses version strings that start with "v" in some places
	//reg := regexp.MustCompile(`\d+?\.\d+?\.\d+`)
	//ver := reg.FindString(tag.TagName)

	// Source code
	fmt.Println(fmt.Sprintf("%s", tag[0].TarballURL))
	fmt.Println("	dir=Editcp")
}

func main() {
	fmt.Println("# https://www.farnsworth.org/dale/codeplug/dmrRadio/downloads")
	fmt.Println("# https://www.farnsworth.org/dale/codeplug/editcp/downloads")
	fmt.Println("# https://www.farnsworth.org/dale/codeplug/editcp")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/codeplug")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/debug")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/dfu")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/dmrRadio")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/docCodeplug")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/docker")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/docs")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/editcp")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/genCodeplugInfo")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/genFileData")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/stdfu")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/ui")
	fmt.Println("# https://github.com/dalefarnsworth-dmr/userdb")

	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/codeplug/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/debug/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/dfu/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/dmrRadio/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/docCodeplug/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/docker/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/docs/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/editcp/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/genCodeplugInfo/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/genFileData/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/stdfu/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/ui/tags")
	dumpOne("https://api.github.com/repos/dalefarnsworth-dmr/userdb/tags")
}
