package main

import (
	"bytes"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"filippo.io/age"
	"github.com/bitfield/script"
	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func AgeKeypair() {
	// XXX FIXME TODO  Check if the files exist first!!!

	// age-keygen 2>/dev/null | tail -1 | (umask 0077 && tee secret_key_age) | age-keygen -y - 2>/dev/null > public_key_age  # generate keypair
	// age-keygen -y secret_key_age > public_key_age  # recover public key
	// (umask 0077 && cat secret_key_age | age --armor --output secret_key_age.age --passphrase)  # add/change password-protection to private key
	// (umask 0077 && age --decrypt --output secret_key_age secret_key_age.age)  # remove password-protection from private key

	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.Fatalf("Failed to generate key pair: %v", err)
	}

	f1, err := os.OpenFile("secret_key_age", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	f1.Write([]byte(fmt.Sprintf("%s\n", identity.String())))
	f1.Close()

	f2, err := os.OpenFile("public_key_age", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	f2.Write([]byte(fmt.Sprintf("%s\n", identity.Recipient().String())))
	f2.Close()
}

func SSHKeypair() {
	// XXX FIXME TODO  Check if the files exist first!!!

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

	_ = ioutil.WriteFile("secret_key_ssh", privateKey, 0600)
	_ = ioutil.WriteFile("public_key_ssh", authorizedKey, 0644)
}

func WireguardKeypair() {
	// XXX FIXME TODO  Check if the files exist first!!!

	// wg genkey | (umask 0077 && tee secret_key_wg) | wg pubkey > public_key_wg  # generate keypair
	// cat secret_key_wg | wg pubkey > public_key_wg  # recover public key

	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		if _, err := os.Stat(fmt.Sprintf("%s/wg", path)); err == nil {
			b1 := new(bytes.Buffer)
			b2 := new(bytes.Buffer)
			_, err := script.Exec("wg genkey").Tee(b1).Exec("wg pubkey").Tee(b2).String()
			if err != nil {
				panic(err)
			}

			f1, err := os.OpenFile("secret_key_wg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			b1.WriteTo(f1)

			f2, err := os.OpenFile("public_key_wg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
			if err != nil {
				panic(err)
			}
			b2.WriteTo(f2)

			break
		}
	}
}

func WireguardPreSharedKey() {
	// XXX FIXME TODO  Check if the files exist first!!!

	// (umask 0077 && wg genpsk > shared_key_wg)  # generate pre-shared key

	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		if _, err := os.Stat(fmt.Sprintf("%s/wg", path)); err == nil {
			b3 := new(bytes.Buffer)
			_, err := script.Exec("wg genpsk").Tee(b3).String()
			if err != nil {
				panic(err)
			}

			f3, err := os.OpenFile("shared_key_wg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			b3.WriteTo(f3)

			break
		}
	}
}

// https://github.com/slaveofcode/age-encryption
// https://github.com/mikalv/anything2ed25519
// https://superuser.com/questions/308126/is-it-possible-to-sign-a-file-using-an-ssh-key
