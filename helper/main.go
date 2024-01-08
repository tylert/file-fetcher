package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func dedupeSlice[T comparable](sliceList []T) []T {
	dedupeMap := make(map[T]struct{})
	list := []T{}

	for _, slice := range sliceList {
		if _, exists := dedupeMap[slice]; !exists {
			dedupeMap[slice] = struct{}{}
			list = append(list, slice)
		}
	}

	return list
}

func main() {
	// Work on entire subdirectories at a time
	dirs, err := script.FindFiles(".").Dirname().Slice()
	if err != nil {
		panic(err)
	}
	dd := dedupeSlice(dirs)

	// Format Go files
	_, err2 := script.Exec("go version").String()
	if err2 != nil {
		fmt.Println("Missing Go binary")
	} else {
		fmt.Println("Formatting Go modules")
		script.Slice(dd).ExecForEach("gofmt -l -w {{ . }}").Stdout()
	}

	// Format Terraform files
	_, err3 := script.Exec("terraform version").String()
	if err3 != nil {
		fmt.Println("Missing Terraform binary")
	} else {
		fmt.Println("Formatting Terraform modules")
		script.Slice(dd).ExecForEach("terraform fmt -list=true -write=true {{ . }}").Stdout()
	}
}
