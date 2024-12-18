package main

import (
	"crypto/rand"
	"encoding/pem"

	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func SshKeypair() (string, string) {
	// ssh-keygen -C '' -N '' -a 16 -f seckey_ssh -t ed25519 ; mv seckey_ssh.pub pubkey_ssh  # generate keypair
	// ssh-keygen -y -f seckey_ssh > pubkey_ssh  # recover public key
	// ssh-keygen -a 512 -p -f seckey_ssh  # add/remove/change password-protection on private key

	mooKey, pooKey, _ := ed25519.GenerateKey(rand.Reader)
	booKey, _ := ssh.NewPublicKey(mooKey)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(pooKey),
	}
	// XXX FIXME TODO  The line-wrapping looks a little funny but it still works fine

	secKey := pem.EncodeToMemory(pemKey)
	pubKey := ssh.MarshalAuthorizedKey(booKey)

	return string(secKey), string(pubKey)
}
