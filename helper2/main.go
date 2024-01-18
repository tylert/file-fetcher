package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

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

	// Slice to remove fill-up zero values which cause a wrong content type detection in the next step
	buff = buff[:n]

	return http.DetectContentType(buff), nil
}

func main() {
	script.FindFiles(".").FilterScan(func(line string, w io.Writer) {
		file, err1 := os.Open(line)
		defer file.Close()
		if err1 != nil {
			fmt.Println(fmt.Sprintf("ERR1 %s", err1))
		}

		// Always returns a valid content-type and "application/octet-stream" if no others seemed to match
		contentType, err2 := Get(file)
		if err2 != nil {
			fmt.Println(fmt.Sprintf("ERR2 %s", err1))
		}

		fmt.Fprintf(w, "%s  %s\n", contentType, line)
	}).Stdout()
}

// XXX FIXME TODO  "sed -i 's/\r//' *" is basically dos2unix

// https://mimesniff.spec.whatwg.org
// https://gist.github.com/rayrutjes/db9b9ea8e02255d62ce2
