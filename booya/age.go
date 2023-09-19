package main

import (
	"fmt"
	"log"
	"os"

	"filippo.io/age"
)

func AgeKeypair(force bool) {
	// age-keygen 2>/dev/null | tail -1 | (umask 0077 && tee secret_key_age) | age-keygen -y - 2>/dev/null > public_key_age  # generate keypair
	// age-keygen -y secret_key_age > public_key_age  # recover public key
	// (umask 0077 && cat secret_key_age | age --armor --output secret_key_age.age --passphrase)  # add/change password-protection to private key
	// (umask 0077 && age --decrypt --output secret_key_age secret_key_age.age)  # remove password-protection from private key

	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.Fatalf("Failed to generate key pair: %v", err)
	}

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	f2, err := os.OpenFile("public_key_age", flags, 0664)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f2.Write([]byte(fmt.Sprintf("%s\n", identity.Recipient().String())))
	f2.Close()

	f1, err := os.OpenFile("secret_key_age", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f1.Write([]byte(fmt.Sprintf("%s\n", identity.String())))
	f1.Close()
}
