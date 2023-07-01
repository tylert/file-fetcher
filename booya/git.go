package main

import (
	"github.com/bitfield/script"
)

func Git() {
	// XXX FIXME TODO  Check if tool is installed first before blindly running it!!!
	script.Echo("You have this version of git installed\n").Stdout()
	script.Exec("git --version").Stdout()
	script.Echo("\n").Stdout()

	script.Echo("You have the following system config set for git\n").Stdout()
	script.Exec("git config --system --list").Stdout()
	script.Echo("\n").Stdout()

	script.Echo("You have the following global config set for git\n").Stdout()
	script.Exec("git config --global --list").Stdout()
	script.Echo("\n").Stdout()

	// XXX FIXME TODO  Force pull.rebase=true to be set if it is missing!!!
	// XXX FIXME TODO  Force user.name, user.email to be set if they are missing!!!
}
