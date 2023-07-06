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
			// Private key
			b1 := new(bytes.Buffer)
			priv, err := script.Exec("age-keygen").Last(1).Tee(b1).String()
			if err != nil {
				panic(err)
			}
			// XXX FIXME TODO  Check if the file exists first!!!
			f1, err := os.OpenFile("age1-sec", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			b1.WriteTo(f1)

			// Public key
			b2 := new(bytes.Buffer)
			_, err2 := script.Echo(priv).Exec("age-keygen -y -").Tee(b2).String()
			if err2 != nil {
				panic(err2)
			}
			// XXX FIXME TODO  Check if the file exists first!!!
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
			// Private key
			b1 := new(bytes.Buffer)
			priv, err := script.Exec("wg genkey").Tee(b1).String()
			if err != nil {
				panic(err)
			}
			// XXX FIXME TODO  Check if the file exists first!!!
			f1, err := os.OpenFile("wg1-sec", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			b1.WriteTo(f1)

			// Public key
			b2 := new(bytes.Buffer)
			_, err2 := script.Echo(priv).Exec("wg pubkey").Tee(b2).String()
			if err2 != nil {
				panic(err2)
			}
			// XXX FIXME TODO  Check if the file exists first!!!
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
