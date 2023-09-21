package main

import (
	"fmt"
	"log"
	"os"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func WireguardKeypair(force bool) {
	// wg genkey | (umask 0077 && tee seckey_wg) | wg pubkey > pubkey_wg  # generate keypair
	// cat seckey_wg | wg pubkey > pubkey_wg  # recover public key

	secKey, _ := wgtypes.GeneratePrivateKey()
	pubKey := secKey.PublicKey()

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	pub, err := os.OpenFile("pubkey_wg", flags, 0644)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer pub.Close()

	sec, err := os.OpenFile("seckey_wg", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer sec.Close()

	sec.Write([]byte(fmt.Sprintf("%s\n", secKey.String())))
	pub.Write([]byte(fmt.Sprintf("%s\n", pubKey.String())))
}

func WireguardPreSharedKey(force bool) {
	// (umask 0077 && wg genpsk > sharedkey_wg)  # generate pre-shared key

	psKey, _ := wgtypes.GenerateKey()

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	prot, err := os.OpenFile("sharedkey_wg", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer prot.Close()

	prot.Write([]byte(fmt.Sprintf("%s\n", psKey.String())))
}
