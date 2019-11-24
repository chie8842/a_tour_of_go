package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (i int, e error) {
	for i, e = 0, nil; i < len(b); i++ {
		b[i] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
