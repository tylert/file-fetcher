package main

import (
	"fmt"
	"log"
	"os"

	"filippo.io/age"
)

func AgeKeypair(force bool) {
	// age-keygen 2>/dev/null | tail -1 | (umask 0077 && tee seckey_age) | age-keygen -y - 2>/dev/null > pubkey_age  # generate keypair
	// age-keygen -y seckey_age > pubkey_age  # recover public key
	// (umask 0077 && cat seckey_age | age --armor --output seckey_age.age --passphrase)  # add/change password-protection to private key
	// (umask 0077 && age --decrypt --output seckey_age seckey_age.age)  # remove password-protection from private key

	identity, _ := age.GenerateX25519Identity()

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	pub, err := os.OpenFile("pubkey_age", flags, 0664)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer pub.Close()

	sec, err := os.OpenFile("seckey_age", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer sec.Close()

	sec.Write([]byte(fmt.Sprintf("%s\n", identity.String())))
	pub.Write([]byte(fmt.Sprintf("%s\n", identity.Recipient().String())))
}
