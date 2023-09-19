package main

import (
	"fmt"
	"log"
	"os"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func WireguardKeypair(force bool) {
	// wg genkey | (umask 0077 && tee secret_key_wg) | wg pubkey > public_key_wg  # generate keypair
	// cat secret_key_wg | wg pubkey > public_key_wg  # recover public key

	priv, _ := wgtypes.GeneratePrivateKey()
	pub := priv.PublicKey()

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	f1, err := os.OpenFile("public_key_wg", flags, 0644)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f1.Write([]byte(fmt.Sprintf("%s\n", pub.String())))
	f1.Close()

	f2, err := os.OpenFile("secret_key_wg", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f2.Write([]byte(fmt.Sprintf("%s\n", priv.String())))
	f2.Close()
}

func WireguardPreSharedKey(force bool) {
	// (umask 0077 && wg genpsk > shared_key_wg)  # generate pre-shared key

	pskey, _ := wgtypes.GenerateKey()

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	f3, err := os.OpenFile("shared_key_wg", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to save file: %v", err)
	}
	f3.Write([]byte(fmt.Sprintf("%s\n", pskey.String())))
	f3.Close()
}
