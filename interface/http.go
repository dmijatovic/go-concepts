package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// define structure to
// implement Write interface
type logWriter struct {
}

// implement standard Write interface to logWriter
func (logWriter) Write(bs []byte) (int, error) {
	// take bs bytes
	// convert bytes to strings
	fmt.Println(string(bs))
	return len(bs), nil
}

// type resp

func mainHTTP() {
	println("Http works!")
	resp, err := http.Get("http://google.com")
	if err != nil {
		println("Err:", err)
		os.Exit(1)
	}
	// this will convert resp type
	// fmt.Println(resp)
	// println(resp)

	// MANUAL APPROACH
	// create byteslice with specific size
	// read function will return required value
	// bs := make([]byte, 4096)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// MAGIC using an method that implements
	// writer interface from Go library
	// io.Copy(os.Stdout, resp.Body)

	// custom Writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
