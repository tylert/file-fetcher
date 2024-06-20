package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"aead.dev/minisign"
)

func MinisignKeypair(force bool) {
	// minisign -fGW -s seckey_ms -p pubkey_ms  # generate keypair
	// minisign -R -s seckey_ms -p pubkey_ms  # recover public key
	// minisign -C -s seckey_ms -p pubkey_ms  # add/remove/change password-protection on private key

	mooKey, pooKey, err := minisign.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("Unable to create key: %v", err)
	}
	secKey, err := minisign.EncryptKey("", pooKey)
	if err != nil {
		log.Fatalf("Unable to encrypt key: %v", err)
	}
	pubKey, _ := mooKey.MarshalText()

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	pub, err := os.OpenFile("pubkey_ms", flags, 0644)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer pub.Close()

	sec, err := os.OpenFile("seckey_ms", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer sec.Close()

	sec.Write([]byte(fmt.Sprintf("%s\n", secKey)))
	pub.Write([]byte(fmt.Sprintf("%s\n", pubKey)))
}
