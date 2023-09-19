package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/ian-kent/envconf"
	"github.com/vharitonsky/iniflags"
)

// Command-line arguments
var (
	aAgeKey      bool
	aForce       bool
	aMinisignKey bool
	aSshKey      bool
	aVersion     bool
	aWgKey       bool
	aWgPsk       bool
)

func init() {
	// Help for command-line arguments
	const (
		sAgeKey      = "Generate an age keypair (default false)"
		sForce       = "Overwrite key files if they exist"
		sMinisignKey = "Generate a minisign keypair (default false)"
		sSshKey      = "Generate an ssh keypair (default false)"
		sVersion     = "Display build version information (default false)"
		sWgKey       = "Generate a wireguard keypair (default false)"
		sWgPsk       = "Generate a wireguard pre-shared key (default false)"
	)

	flag.BoolVar(&aAgeKey, "age", FromEnvP("MEH_AGE", false).(bool), sAgeKey)
	flag.BoolVar(&aAgeKey, "a", FromEnvP("MEH_AGE", false).(bool), sAgeKey)
	flag.BoolVar(&aForce, "force", FromEnvP("MEH_FORCE", false).(bool), sForce)
	flag.BoolVar(&aForce, "f", FromEnvP("MEH_FORCE", false).(bool), sForce)
	flag.BoolVar(&aMinisignKey, "minisign", FromEnvP("MEH_MINISIGN", false).(bool), sMinisignKey)
	flag.BoolVar(&aMinisignKey, "m", FromEnvP("MEH_MINISIGN", false).(bool), sMinisignKey)
	flag.BoolVar(&aSshKey, "ssh", FromEnvP("MEH_SSH", false).(bool), sSshKey)
	flag.BoolVar(&aSshKey, "s", FromEnvP("MEH_SSH", false).(bool), sSshKey)
	flag.BoolVar(&aVersion, "version", FromEnvP("MEH_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("MEH_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aWgKey, "wg", FromEnvP("MEH_WG", false).(bool), sWgKey)
	flag.BoolVar(&aWgKey, "w", FromEnvP("MEH_WG", false).(bool), sWgKey)
	flag.BoolVar(&aWgPsk, "wgpsk", FromEnvP("MEH_WGPSK", false).(bool), sWgPsk)
	flag.BoolVar(&aWgPsk, "p", FromEnvP("MEH_WGPSK", false).(bool), sWgPsk)
	iniflags.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
