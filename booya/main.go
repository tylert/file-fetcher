package main

import (
// "github.com/bitfield/script"
)

func main() {
	Python()
	Pyenv()
	Git()
	//script.Get("https://wttr.in/Ottawa?m&format=3").Stdout()
	//script.Get("https://wttr.in/Ottawa?u&format=3").Stdout()
}

// Sec-CH-UA                   => XXX
// Sec-CH-UA-Arch              => "x86", "arm"
// Sec-CH-UA-Bitness           => 64
// Sec-CH-UA-Form-Factor       => ""
// Sec-CH-UA-Full-Version      => "1.0.0"
// Sec-CH-UA-Full-Version-List => XXX
// Sec-CH-UA-Mobile            => ?0
// Sec-CH-UA-Model             => XXX
// Sec-CH-UA-Platform          => "Linux", "macOS", "Windows"
// Sec-CH-UA-Platform-Version  => "", "13.4.1", "11"
// Sec-CH-UA-WoW64             => ?1, ?0

// https://wicg.github.io/ua-client-hints/#sec-ch-ua-platform-version

// These might seem a bit silly to ask since you already had to compile this for their OS/CPU...
// (work on both macOS and Linux)
// getconf LONG_BIT
// uname -m, uname -p, uname -i, uname -s, uname -o
// arch
// [ $((0xffffffff)) -eq -1 ] && echo 32 || echo 64
