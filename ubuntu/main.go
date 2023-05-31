package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"strings"
)

type Product struct {
	Format string `json:"format"`
	Index  struct {
		ComUbuntuReleasesUbuntu struct {
			Datatype string   `json:"datatype"`
			Format   string   `json:"format"`
			Path     string   `json:"path"`
			Products []string `json:"products"`
			Updated  string   `json:"updated"`
		} `json:"com.ubuntu.releases:ubuntu"`
		ComUbuntuReleasesUbuntuServer struct {
			Datatype string   `json:"datatype"`
			Format   string   `json:"format"`
			Path     string   `json:"path"`
			Products []string `json:"products"`
			Updated  string   `json:"updated"`
		} `json:"com.ubuntu.releases:ubuntu-server"`
	} `json:"index"`
	Updated struct {
		Datatype string `json:"datatype"`
		Updated  string `json:"updated"`
	} `json:"updated"`
}

func main() {
	res, err := http.Get("http://releases.ubuntu.com/streams/v1/index.json")
	//res, err := http.Get("http://releases.ubuntu.com/streams/v1/com.ubuntu.releases:ubuntu-server.json")
	//res, err := http.Get("http://releases.ubuntu.com/streams/v1/com.ubuntu.releases:ubuntu.json")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	var pro Product
	err = json.NewDecoder(res.Body).Decode(&pro)
	if err != nil {
		log.Fatalf("Kaboom!")
	}

	fmt.Println("# http://releases.ubuntu.com/streams/v1/")
	fmt.Println("# https://releases.ubuntu.com")
	fmt.Println("# https://cdimage.ubuntu.com")
	fmt.Println("# https://ubuntu.com/download/raspberry-pi")

	// Live Server
	for i := 0; i < len(pro.Index.ComUbuntuReleasesUbuntuServer.Products); i++ {
		fmt.Println(pro.Index.ComUbuntuReleasesUbuntuServer.Products[i])
	}

	// Desktop
	for i := 0; i < len(pro.Index.ComUbuntuReleasesUbuntu.Products); i++ {
		fmt.Println(pro.Index.ComUbuntuReleasesUbuntu.Products[i])
	}
}
