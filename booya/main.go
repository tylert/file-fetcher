package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if aVersion {
		fmt.Println(GetVersion())
	}

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !aForce {
		flags |= os.O_EXCL
	}

	if aAgeKey {
		sec, pub := AgeKeypair()

		secf, err := os.OpenFile("seckey_age", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		pubf, err := os.OpenFile("pubkey_age", flags, 0664)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s\n", pub)))
	}

	if aMinisignKey {
		sec, pub := MinisignKeypair()

		secf, err := os.OpenFile("seckey_ms", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		pubf, err := os.OpenFile("pubkey_ms", flags, 0644)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s\n", pub)))
	}

	if aNncpKeys {
		sec, pub := NncpConfigData()

		secf, err := os.OpenFile("seckeys_nncp", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		pubf, err := os.OpenFile("pubkeys_nncp", flags, 0644)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s\n", pub)))
	}

	if aSshKey {
		sec, pub := SshKeypair()

		secf, err := os.OpenFile("seckey_ssh", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		pubf, err := os.OpenFile("pubkey_ssh", flags, 0664)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf.Write([]byte(fmt.Sprintf("%s", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s", pub)))
	}

	if aWgKey {
		sec, pub := WireguardKeypair()

		secf, err := os.OpenFile("seckey_wg", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		pubf, err := os.OpenFile("pubkey_wg", flags, 0644)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s\n", pub)))
	}

	if aWgPsk {
		sec := WireguardPreSharedKey()

		secf, err := os.OpenFile("secpsk_wg", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
	}
}
