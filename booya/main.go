package main

import (
	"log"
	"os"

	"github.com/bitfield/script"
)

func main() {
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

// https://bitfieldconsulting.com/golang/scripting
// https://github.com/bitfield/script
// https://pkg.go.dev/github.com/bitfield/script
