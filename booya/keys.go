package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/bitfield/script"
)

func AgeKeypair() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(fmt.Sprintf("%s/age-keygen", paths[i])); err == nil {
			// XXX FIXME TODO  Check if the files exist first!!!
			b1 := new(bytes.Buffer)
			b2 := new(bytes.Buffer)
			_, err := script.Exec("age-keygen").Last(1).Tee(b1).Exec("age-keygen -y -").Tee(b2).String()
			if err != nil {
				panic(err)
			}

			// Write out the private key file
			f1, err := os.OpenFile("age1-sec", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			b1.WriteTo(f1)

			// Write out the public key file
			f2, err := os.OpenFile("age1-pub", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
			if err != nil {
				panic(err)
			}
			b2.WriteTo(f2)
		}
	}
}

func WireguardKeypair() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(fmt.Sprintf("%s/wg", paths[i])); err == nil {
			// XXX FIXME TODO  Check if the files exist first!!!
			b1 := new(bytes.Buffer)
			b2 := new(bytes.Buffer)
			_, err := script.Exec("wg genkey").Tee(b1).Exec("wg pubkey").Tee(b2).String()
			if err != nil {
				panic(err)
			}

			// Write out the private key file
			f1, err := os.OpenFile("wg1-sec", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			b1.WriteTo(f1)

			// Write out the public key file
			f2, err := os.OpenFile("wg1-pub", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
			if err != nil {
				panic(err)
			}
			b2.WriteTo(f2)
		}
	}
}

func SSHKeypair() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(fmt.Sprintf("%s/ssh-keygen", paths[i])); err == nil {
			_, err := script.Exec("ssh-keygen -N '' -C '' -f ssh1 -a 65535 -t ed25519").String()
			if err != nil {
				panic(err)
			}
		}
	}
}
