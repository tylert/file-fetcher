package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"aead.dev/minisign"
)

func MinisignKeypair(force bool) {
	// minisign -GWf -s secret_key_minisign -p public_key_minisign  # generate keypair
	// minisign -R -s secret_key_minisign -p public_key_minisign  # recover public key
	// minisign -C -s secret_key_minisign -p public_key_minisign  # add/remove/change password-protection on private key

	public, private, err := minisign.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	f1, err := os.OpenFile("public_key_minisign", flags, 0644)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f1.Write([]byte(fmt.Sprintf("%s\n", public)))
	f1.Close()

	f2, err := os.OpenFile("secret_key_minisign", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f2.Write([]byte(fmt.Sprintf("%s\n", private)))
	f2.Close()
}
