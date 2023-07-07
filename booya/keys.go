package main

import (
	"bytes"
	"crypto/rand"
	"encoding/pem"
	"io/ioutil"
	"os"

	"github.com/bitfield/script"
	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func AgeKeypair() {
	// XXX FIXME TODO  Check if the files exist first!!!
	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	// age-keygen 2>/dev/null | tail -1 | tee priv | age-keygen -y - 2>/dev/null > pub
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

func WireguardKeypair() {
	// XXX FIXME TODO  Check if the files exist first!!!
	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	// wg genkey | tee priv | wg pubkey > pub
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

func SSHKeypair() {
	// XXX FIXME TODO  Check if the files exist first!!!
	// ssh-keygen -C '' -N '' -a 16 -f ssh1-sec -t ed25519 ; mv ssh1-sec.pub ssh1-pub
	pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
	publicKey, _ := ssh.NewPublicKey(pubKey)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privKey), // <- marshals ed25519 correctly
	}

	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	_ = ioutil.WriteFile("ssh1-sec", privateKey, 0600)
	_ = ioutil.WriteFile("ssh1-pub", authorizedKey, 0644)
}
