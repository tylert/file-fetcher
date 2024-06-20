package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"gopkg.in/ini.v1"
)

// Command-line arguments
var (
	aAgeKey      bool
	aForce       bool
	aMinisignKey bool
	aNncpKeys    bool
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
		uNncpKeys    = "Generate a set of nncp node keypairs (default false)"
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
	flag.BoolVar(&aNncpKeys, "nncp", FromEnvP("MEH_NNCP", false).(bool), uNncpKeys)
	flag.BoolVar(&aNncpKeys, "n", FromEnvP("MEH_NNCP", false).(bool), uNncpKeys)
	flag.BoolVar(&aSshKey, "ssh", FromEnvP("MEH_SSH", false).(bool), uSshKey)
	flag.BoolVar(&aSshKey, "s", FromEnvP("MEH_SSH", false).(bool), uSshKey)
	flag.BoolVar(&aVersion, "version", FromEnvP("MEH_VERSION", false).(bool), uVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("MEH_VERSION", false).(bool), uVersion)
	flag.BoolVar(&aWgKey, "wg", FromEnvP("MEH_WG", false).(bool), uWgKey)
	flag.BoolVar(&aWgKey, "w", FromEnvP("MEH_WG", false).(bool), uWgKey)
	flag.BoolVar(&aWgPsk, "wgpsk", FromEnvP("MEH_WGPSK", false).(bool), uWgPsk)
	flag.BoolVar(&aWgPsk, "p", FromEnvP("MEH_WGPSK", false).(bool), uWgPsk)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		// flag.VisitAll(func(f *flag.Flag) {
		//   fmt.Fprintf(os.Stderr, "%v %v %v\n", f.Name, f.Value, f.Usage)
		// })
	}

	// FlagSet for sub-commands???
	// https://digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go

	// Attempt to gracefully load things from a known config file location
	// cfg := ini.Empty()
	// home, _ := os.UserHomeDir()
	// cfg, err := ini.LooseLoad(fmt.Sprintf("%s/.config/moo/defaults", home))
	// https://ini.unknwon.io/docs

	flag.Parse()
	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}

var (
	// ErrUnsupportedType is returned if the type passed in is unsupported
	ErrUnsupportedType = errors.New("Unsupported type")
)

// FromEnvP is the same as FromEnv, but panics on error
func FromEnvP(env string, value interface{}) interface{} {
	ev, err := FromEnv(env, value)
	if err != nil {
		panic(err)
	}
	return ev
}

// FromEnv returns the environment variable specified by env
// using the type of value
func FromEnv(env string, value interface{}) (interface{}, error) {
	envs := os.Environ()
	found := false
	for _, e := range envs {
		if strings.HasPrefix(e, env+"=") {
			found = true
			break
		}
	}

	if !found {
		return value, nil
	}

	ev := os.Getenv(env)

	switch value.(type) {
	case string:
		vt := interface{}(ev)
		return vt, nil
	case int:
		i, e := strconv.ParseInt(ev, 10, 64)
		return int(i), e
	case int8:
		i, e := strconv.ParseInt(ev, 10, 8)
		return int8(i), e
	case int16:
		i, e := strconv.ParseInt(ev, 10, 16)
		return int16(i), e
	case int32:
		i, e := strconv.ParseInt(ev, 10, 32)
		return int32(i), e
	case int64:
		i, e := strconv.ParseInt(ev, 10, 64)
		return i, e
	case uint:
		i, e := strconv.ParseUint(ev, 10, 64)
		return uint(i), e
	case uint8:
		i, e := strconv.ParseUint(ev, 10, 8)
		return uint8(i), e
	case uint16:
		i, e := strconv.ParseUint(ev, 10, 16)
		return uint16(i), e
	case uint32:
		i, e := strconv.ParseUint(ev, 10, 32)
		return uint32(i), e
	case uint64:
		i, e := strconv.ParseUint(ev, 10, 64)
		return i, e
	case float32:
		i, e := strconv.ParseFloat(ev, 32)
		return float32(i), e
	case float64:
		i, e := strconv.ParseFloat(ev, 64)
		return float64(i), e
	case bool:
		i, e := strconv.ParseBool(ev)
		return i, e
	default:
		return value, ErrUnsupportedType
	}
}
