package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bitfield/script"
)

func Python() {
	// XXX FIXME TODO  Also try to find things named:  python2, python3, python37, python3.9, etc.
	paths := strings.Split(os.Getenv("PATH"), ":")
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(fmt.Sprintf("%s/python", paths[i])); err == nil {
			script.Echo(fmt.Sprintf("This is the version of python installed at %s/python:\n", paths[i])).Stdout()
			script.Exec(fmt.Sprintf("%s/python --version", paths[i])).Stdout()
			script.Echo("\n").Stdout()

			script.Echo("This instance of python has the following 'site' module:\n").Stdout()
			script.Exec(fmt.Sprintf("%s/python -m site", paths[i])).Stdout()
			script.Echo("\n").Stdout()
		}
	}

	// XXX FIXME TODO  Nag them if it isn't the latest???  Or at least set a sensible lower-bound!!!
}

func Pyenv() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(fmt.Sprintf("%s/pyenv", paths[i])); err == nil {
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
	}

	// XXX FIXME TODO  Nag them if it isn't the latest???  Or at least set a sensible lower-bound!!!
}
