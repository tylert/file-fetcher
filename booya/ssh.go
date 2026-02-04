package main

import (
	"crypto/rand"
	"encoding/pem"

	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func SshKeypair() (string, string) {
	// ssh-keygen -f seckey_ssh -t ed25519 -q -C '' -N '' ; mv seckey_ssh.pub pubkey_ssh  # generate keypair
	// ssh-keygen -f seckey_ssh -y > pubkey_ssh  # recover public key
	// ssh-keygen -f seckey_ssh -c  # change comment (repack private key)
	// ssh-keygen -f seckey_ssh -p  # change password (repack private key)
	// dropbearconvert openssh dropbear seckey_ssh sec2key_ssh  # convert key for dropbear (unencrypted only)
	// dropbearconvert dropbear openssh sec2key_ssh seckey_ssh  # convert key for openssh (unencrypted only)

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
