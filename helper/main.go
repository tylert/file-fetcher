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
	dirs, err1 := script.FindFiles(".").Dirname().Slice()
	if err1 != nil {
		panic(err1)
	}
	dd := dedupeSlice(dirs)

	// Go files
	_, err2 := script.Exec("go version").String()
	if err2 != nil {
		fmt.Println("Missing Go binary")
	} else {
		fmt.Println("Formatting Go modules")
		script.Slice(dd).ExecForEach("gofmt -l -w {{ . }}").Stdout()
	}

	// OpenTofu files
	_, err3 := script.Exec("tofu version").String()
	if err3 != nil {
		fmt.Println("Missing OpenTofu binary")
	} else {
		fmt.Println("Formatting OpenTofu modules")
		script.Slice(dd).ExecForEach("tofu fmt -list=true -write=true {{ . }}").Stdout()
	}

	_, err4 := script.Exec("tflint --version").String()
	if err4 != nil {
		fmt.Println("Missing TFLint binary")
	} else {
		fmt.Println("Linting OpenTofu modules")
		script.Slice(dd).ExecForEach("tflint --chdir={{ . }}").Stdout()
	}
}
