package iotools

import (
	"io"
	"io/ioutil"
	"os"
)

// ReadStdinOrFile reads from stdin if it is not empty. Else, it reads from the
// given file. In either case, it returns the entire data read, and error if
// any.
func ReadStdinOrFile(f string) ([]byte, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	var reader io.ReadCloser
	if stat.Size() > 0 {
		reader = os.Stdin
	} else {
		var err error
		reader, err = os.Open(f)
		if err != nil {
			panic(err)
		}
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
