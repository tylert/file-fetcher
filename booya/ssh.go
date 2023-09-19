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
	// ssh-keygen -C '' -N '' -a 16 -f secret_key_ssh -t ed25519 ; mv secret_key_ssh.pub public_key_ssh  # generate keypair
	// ssh-keygen -y -f secret_key_ssh > public_key_ssh  # recover public key
	// ssh-keygen -a 512 -p -f secret_key_ssh  # add/remove/change password-protection on private key

	pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
	publicKey, _ := ssh.NewPublicKey(pubKey)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privKey), // marshals ed25519 correctly
	}
	// XXX FIXME TODO  The line-wrapping looks a little funny but it still works fine

	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	f1, err := os.OpenFile("public_key_ssh", flags, 0664)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f1.Write([]byte(fmt.Sprintf("%s", authorizedKey)))
	f1.Close()

	f2, err := os.OpenFile("secret_key_ssh", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f2.Write([]byte(fmt.Sprintf("%s", privateKey)))
	f2.Close()
}

// https://github.com/mikalv/anything2ed25519
// https://superuser.com/questions/308126/is-it-possible-to-sign-a-file-using-an-ssh-key
