package main

import (
	"fmt"
)

func main() {
	if aVersion {
		fmt.Println(GetVersion())
	}

	if aAgeKey {
		AgeKeypair(aForce)
	}

	if aMinisignKey {
		MinisignKeypair(aForce)
	}

	if aNncpKeys {
		NncpConfigData(aForce)
	}

	if aSshKey {
		SshKeypair(aForce)
	}

	if aWgKey {
		WireguardKeypair(aForce)
	}

	if aWgPsk {
		WireguardPreSharedKey(aForce)
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
