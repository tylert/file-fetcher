package main

import (
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func SSHKeypair(force bool) {
	// ssh-keygen -C '' -N '' -a 16 -f seckey_ssh -t ed25519 ; mv seckey_ssh.pub pubkey_ssh  # generate keypair
	// ssh-keygen -y -f seckey_ssh > pubkey_ssh  # recover public key
	// ssh-keygen -a 512 -p -f seckey_ssh  # add/remove/change password-protection on private key

	mooKey, pooKey, _ := ed25519.GenerateKey(rand.Reader)
	booKey, _ := ssh.NewPublicKey(mooKey)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(pooKey), // marshals ed25519 correctly
	}
	// XXX FIXME TODO  The line-wrapping looks a little funny but it still works fine

	secKey := pem.EncodeToMemory(pemKey)
	pubKey := ssh.MarshalAuthorizedKey(booKey)

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	pub, err := os.OpenFile("pubkey_ssh", flags, 0664)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer pub.Close()

	sec, err := os.OpenFile("seckey_ssh", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer sec.Close()

	sec.Write([]byte(fmt.Sprintf("%s", secKey)))
	pub.Write([]byte(fmt.Sprintf("%s", pubKey)))
}

// https://github.com/mikalv/anything2ed25519
// https://superuser.com/questions/308126/is-it-possible-to-sign-a-file-using-an-ssh-key
