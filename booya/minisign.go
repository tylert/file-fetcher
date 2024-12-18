package main

import (
	"crypto/rand"
	"log"

	"aead.dev/minisign"
)

func MinisignKeypair() (string, string) {
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

	return string(secKey), string(pubKey)
}
