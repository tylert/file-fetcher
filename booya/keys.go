package main

import (
	"bytes"
	"crypto/rand"
	"encoding/pem"
	"io/ioutil"
	"os"

	"github.com/bitfield/script"
	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func AgeKeypair() {
	// XXX FIXME TODO  Check if the tools/files exist first!!!

	// age-keygen 2>/dev/null | tail -1 | tee secret_key_age | age-keygen -y - 2>/dev/null > public_key_age  # generate keypair
	// cat secret_key_file | age -p > meh && mv meh secret_key_file  # set/change passphrase
	// age-keygen -y secret_key_age > public_key_age  # recover public key

	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	_, err := script.Exec("age-keygen").Last(1).Tee(b1).Exec("age-keygen -y -").Tee(b2).String()
	if err != nil {
		panic(err)
	}

	f1, err := os.OpenFile("secret_key_age", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	b1.WriteTo(f1)

	f2, err := os.OpenFile("public_key_age", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	b2.WriteTo(f2)
}

func SSHKeypair() {
	// XXX FIXME TODO  Check if the files exist first!!!

	// ssh-keygen -C '' -N '' -a 16 -f secret_key_ssh -t ed25519 ; mv secret_key_ssh.pub public_key_ssh  # generate keypair
	// ssh-keygen -a 512 -p -f secret_key_ssh  # set/change passphrase
	// ssh-keygen -y -f secret_key_ssh > public_key_ssh  # recover public key

	pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
	publicKey, _ := ssh.NewPublicKey(pubKey)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privKey), // marshals ed25519 correctly
	}

	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	_ = ioutil.WriteFile("secret_key_ssh", privateKey, 0600)
	_ = ioutil.WriteFile("public_key_ssh", authorizedKey, 0644)
}

func WireguardKeypair() {
	// XXX FIXME TODO  Check if the tools/files exist first!!!

	// wg genkey | tee secret_key_wg | wg pubkey > public_key_wg  # generate keypair
	// cat secret_key_wg | wg pubkey > public_key_wg  # recover public key

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
}

func WireguardPreSharedKey() {
	// XXX FIXME TODO  Check if the tools/files exist first!!!

	// wg genpsk > shared_key_wg  # generate pre-shared key

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
}
