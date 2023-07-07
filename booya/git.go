package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bitfield/script"
)

func Git() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(fmt.Sprintf("%s/python", paths[i])); err == nil {
			script.Echo("You have this version of git installed\n").Stdout()
			script.Exec("git --version").Stdout()
			script.Echo("\n").Stdout()

			script.Echo("You have the following system config set for git\n").Stdout()
			script.Exec("git config --system --list").Stdout()
			script.Echo("\n").Stdout()

			script.Echo("You have the following global config set for git\n").Stdout()
			script.Exec("git config --global --list").Stdout()
			script.Echo("\n").Stdout()
		}
	}

	// XXX FIXME TODO  Nag them if it isn't the latest???  Or at least set a sensible lower-bound!!!
	// XXX FIXME TODO  Force pull.rebase=true to be set if it is missing!!!
	// XXX FIXME TODO  Force user.name, user.email to be set if they are missing!!!
}
