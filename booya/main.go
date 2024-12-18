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

		pubf, err := os.OpenFile("pubkey_age", flags, 0664)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf, err := os.OpenFile("seckey_age", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s\n", pub)))
	}

	if aMinisignKey {
		sec, pub := MinisignKeypair()

		pubf, err := os.OpenFile("pubkey_ms", flags, 0644)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf, err := os.OpenFile("seckey_ms", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s\n", pub)))
	}

	if aNncpKeys {
		NncpConfigData(aForce)
	}

	if aSshKey {
		sec, pub := SshKeypair()

		pubf, err := os.OpenFile("pubkey_ssh", flags, 0664)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf, err := os.OpenFile("seckey_ssh", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		secf.Write([]byte(fmt.Sprintf("%s", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s", pub)))
	}

	if aWgKey {
		sec, pub := WireguardKeypair()

		pubf, err := os.OpenFile("pubkey_wg", flags, 0644)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer pubf.Close()

		secf, err := os.OpenFile("seckey_wg", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer secf.Close()

		secf.Write([]byte(fmt.Sprintf("%s\n", sec)))
		pubf.Write([]byte(fmt.Sprintf("%s\n", pub)))
	}

	if aWgPsk {
		prot := WireguardPreSharedKey()

		protf, err := os.OpenFile("secpsk_wg", flags, 0600)
		if err != nil {
			log.Fatalf("Unable to open file: %v", err)
		}
		defer protf.Close()

		protf.Write([]byte(fmt.Sprintf("%s\n", prot)))
	}
}

// Sec-CH-UA                   => meh...
// Sec-CH-UA-Arch              => "x86", "ARM"
// Sec-CH-UA-Bitness           => "64", "32"
// Sec-CH-UA-Form-Factor       => "Desktop"
// Sec-CH-UA-Full-Version-List => meh...
// Sec-CH-UA-Mobile            => ?0
// Sec-CH-UA-Model             => ""
// Sec-CH-UA-Platform          => "Linux", "macOS", "Windows"
// Sec-CH-UA-Platform-Version  => "", "14.5", "11"

// https://wicg.github.io/ua-client-hints/#sec-ch-ua-platform-version

// It might seem a bit silly to ask for these since you already had to compile this for their OS/CPU...
// (most work on both Linux and macOS)
// [ $((0xffffffff)) -eq -1 ] && echo 32 || echo 64  # "64", "32"
// getconf LONG_BIT  # "64", "32"
// uname -m  # "x86_64", "arm64", etc.
// uname -s  # "Linux", "Darwin", etc.
// uname -o  # "GNU/Linux", "Darwin", etc.
// uname -p  # "unknown", "arm", etc.
// uname -i  # "unknown", illegal option, etc.
// arch  # doesn't work on some Linux distros
