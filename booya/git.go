package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bitfield/script"
)

func Git() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for i := 0; i < len(paths); i++ {
		if _, err := os.Stat(fmt.Sprintf("%s/git", paths[i])); err == nil {
			script.Echo(fmt.Sprintf("You have this version of git installed in %s:\n", paths[i])).Stdout()
			script.Exec(fmt.Sprintf("%s/git --version", paths[i])).Stdout()
			script.Echo("\n").Stdout()

			script.Echo("You have the following system config set for git:\n").Stdout()
			script.Exec(fmt.Sprintf("%s/git config --system --list", paths[i])).Stdout()
			script.Echo("\n").Stdout()
		}
	}
}

func GitConfig() {
	script.Echo("You have the following global config set for git:\n").Stdout()
	script.Exec("git config --global --list").Stdout()
	script.Echo("\n").Stdout()

	num2, err2 := script.Exec("git config --global --list").Match("user.name=").CountLines()
	if err2 != nil {
		log.Fatal(err2)
	}
	if num2 != 1 {
		fmt.Println("The git setting 'user.name' is a required field.")
	}

	num3, err3 := script.Exec("git config --global --list").Match("user.email=").CountLines()
	if err3 != nil {
		log.Fatal(err3)
	}
	if num3 != 1 {
		fmt.Println("The git setting 'user.email' is a required field.")
	}

	num4, err4 := script.Exec("git config --global --list").Match("pull.rebase=true").CountLines()
	if err4 != nil {
		log.Fatal(err4)
	}
	if num4 != 1 {
		fmt.Println("The git setting 'pull.rebase=true' is a required field.")
	}
}
