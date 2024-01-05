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
	dirs, err := script.FindFiles(".").Dirname().Slice()
	if err != nil {
		panic(err)
	}
	dd := dedupeSlice(dirs)

	fmt.Println("Formatting Go modules")
	script.Slice(dd).ExecForEach("gofmt -l -w {{ . }}").Stdout()
	// fmt.Println("Formatting Terraform modules")
	// script.Slice(dd).ExecForEach("terraform fmt -list=true -write=true {{ . }}").Stdout()
}
