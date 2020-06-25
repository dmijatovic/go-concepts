package main

import (
	"fmt"
	"io"
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

func (logWriter) Read(bs []byte) (int, error) {
	// take bs bytes
	// convert bytes to strings
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	args := os.Args
	println("It works")
	fileName := args[1]
	fmt.Println("Filename:", fileName)
	if fileName != "" {
		f, e := os.Open(fileName)
		// MANUALY
		//create read buffer
		// b := make([]byte, 1024)
		// n1, e := f.Read(b)
		if e != nil {
			fmt.Println("Error:", e)
			os.Exit(1)
		}
		// // f.Read()
		// fmt.Println(string(b[:n1]))
		// lw := logWriter{}
		// io.Copy(lw, f.Read())
		// pass
		io.Copy(os.Stdout, f)
	}
}
