package main

import (
	"log"
	"os"

	"github.com/bitfield/script"
)

func git() {
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

func foop() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//home, _ := os.UserHomeDir()
	//err := os.Chdir(filepath.Join(home, "git"))
	//if err != nil {
	//	panic(err)
	//}

	script.FindFiles(cwd).Stdout()
}

func main() {
	git()
}

// https://bitfieldconsulting.com/golang/scripting
// https://github.com/bitfield/script
// https://pkg.go.dev/github.com/bitfield/script
