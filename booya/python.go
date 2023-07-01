package main

import (
	"regexp"

	"github.com/bitfield/script"
)

func Python() {
	// XXX FIXME TODO  Check if tool is installed first before blindly running it!!!
	script.Echo("Moo:\n").Stdout()
	script.Exec("python --version").Stdout()
	script.Echo("\n").Stdout()

	script.Echo("Moo:\n").Stdout()
	script.Exec("python -m site").Stdout()
	script.Echo("\n").Stdout()
}

func Pyenv() {
	// XXX FIXME TODO  Check if tool is installed first before blindly running it!!!
	script.Echo("You have this version of pyenv installed:\n").Stdout()
	script.Exec("pyenv --version").Stdout()
	script.Echo("\n").Stdout()

	// pyenv install --list | grep -E '\s3' | tail
	script.Echo("These are the newest versions of python available for your pyenv:\n").Stdout()
	script.Exec("pyenv install --list").MatchRegexp(regexp.MustCompile(`\s3`)).Last(10).Stdout()
	script.Echo("\n").Stdout()

	script.Echo("You have these versions of python already installed in your pyenv:\n").Stdout()
	script.Exec("pyenv versions").Stdout()
	script.Echo("\n").Stdout()
}
