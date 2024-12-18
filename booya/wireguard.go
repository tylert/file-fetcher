package main

import (
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func WireguardKeypair() (string, string) {
	// wg genkey | (umask 0077 && tee seckey_wg) | wg pubkey > pubkey_wg  # generate keypair
	// cat seckey_wg | wg pubkey > pubkey_wg  # recover public key

	secKey, _ := wgtypes.GeneratePrivateKey()
	pubKey := secKey.PublicKey()

	return secKey.String(), pubKey.String()
}

func WireguardPreSharedKey() string {
	// (umask 0077 && wg genpsk > secpsk_wg)  # generate pre-shared key

	psKey, _ := wgtypes.GenerateKey()

	return psKey.String()
}
