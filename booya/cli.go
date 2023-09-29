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
	// Usage for command-line arguments
	const (
		uAgeKey      = "Generate an age keypair (default false)"
		uForce       = "Overwrite key files if they exist"
		uMinisignKey = "Generate a minisign keypair (default false)"
		uSshKey      = "Generate an ssh keypair (default false)"
		uVersion     = "Display build version information (default false)"
		uWgKey       = "Generate a wireguard keypair (default false)"
		uWgPsk       = "Generate a wireguard pre-shared key (default false)"
	)

	flag.BoolVar(&aAgeKey, "age", FromEnvP("MEH_AGE", false).(bool), uAgeKey)
	flag.BoolVar(&aAgeKey, "a", FromEnvP("MEH_AGE", false).(bool), uAgeKey)
	flag.BoolVar(&aForce, "force", FromEnvP("MEH_FORCE", false).(bool), uForce)
	flag.BoolVar(&aForce, "f", FromEnvP("MEH_FORCE", false).(bool), uForce)
	flag.BoolVar(&aMinisignKey, "minisign", FromEnvP("MEH_MINISIGN", false).(bool), uMinisignKey)
	flag.BoolVar(&aMinisignKey, "m", FromEnvP("MEH_MINISIGN", false).(bool), uMinisignKey)
	flag.BoolVar(&aSshKey, "ssh", FromEnvP("MEH_SSH", false).(bool), uSshKey)
	flag.BoolVar(&aSshKey, "s", FromEnvP("MEH_SSH", false).(bool), uSshKey)
	flag.BoolVar(&aVersion, "version", FromEnvP("MEH_VERSION", false).(bool), uVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("MEH_VERSION", false).(bool), uVersion)
	flag.BoolVar(&aWgKey, "wg", FromEnvP("MEH_WG", false).(bool), uWgKey)
	flag.BoolVar(&aWgKey, "w", FromEnvP("MEH_WG", false).(bool), uWgKey)
	flag.BoolVar(&aWgPsk, "wgpsk", FromEnvP("MEH_WGPSK", false).(bool), uWgPsk)
	flag.BoolVar(&aWgPsk, "p", FromEnvP("MEH_WGPSK", false).(bool), uWgPsk)
	iniflags.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
