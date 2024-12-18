package main

import (
	"filippo.io/age"
)

func AgeKeypair() (string, string) {
	// age-keygen 2>/dev/null | tail -1 | (umask 0077 && tee seckey_age) | age-keygen -y - 2>/dev/null > pubkey_age  # generate keypair
	// age-keygen -y seckey_age > pubkey_age  # recover public key
	// (umask 0077 && cat seckey_age | age --armor --output seckey_age.age --passphrase)  # add/change password-protection to private key
	// (umask 0077 && age --decrypt --output seckey_age seckey_age.age)  # remove password-protection from private key

	identity, _ := age.GenerateX25519Identity()

	return identity.String(), identity.Recipient().String()
}
