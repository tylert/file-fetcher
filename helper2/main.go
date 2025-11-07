package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

func Get(seeker io.ReadSeeker) (string, error) {
	// At most the first 512 bytes of data are used:
	// https://golang.org/src/net/http/sniff.go?s=646:688#L11
	buff := make([]byte, 512)

	_, err := seeker.Seek(0, io.SeekStart)
	if err != nil {
		return "UNKNOWN", err
	}

	n, err := seeker.Read(buff)
	if err != nil && err != io.EOF {
		return "UNKNOWN", err
	}

	// Slice to fill-up zero values which cause a wrong content type detection
	buff = buff[:n]

	return http.DetectContentType(buff), nil
}

func main() {
	nhd, err0 := script.FindFiles(".").FilterScan(func(line string, w io.Writer) {
		file, err1 := os.Open(line)
		defer file.Close()
		if err1 != nil {
			fmt.Println(fmt.Sprintf("ERR1 %s", err1))
		}

		// "application/octet-stream" if no others seemed to match
		contentType, err2 := Get(file)
		if err2 != nil {
			fmt.Println(fmt.Sprintf("ERR2 %s", err2))
		}

		// Only show things that are NOT a hotdog
		match, _ := regexp.MatchString("^text", contentType)
		if match {
			fmt.Fprintf(w, "%s\n", line)
		}
	}).Slice()
	if err0 != nil {
		panic(err0)
	}

	reg := regexp.MustCompile("\r")
	_, err3 := script.Slice(nhd).ReplaceRegexp(reg, "replacement").Stdout()
	if err3 != nil {
		panic(err3)
	}
}

// XXX FIXME TODO  "sed -i 's/\r//' *" is basically dos2unix

// https://mimesniff.spec.whatwg.org
// https://gist.github.com/rayrutjes/db9b9ea8e02255d62ce2
// https://github.com/xyproto/mime
