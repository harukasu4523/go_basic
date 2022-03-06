package main

import (
	"io"
	"os"
)

func main() {
	copy := "test1.go"
	Newfile, err := os.Create("copy.go")
	if err != nil {
		panic(err)
	}
	Oldfile, err := os.Open(copy)
	if err != nil {
		panic(err)
	}
	defer Oldfile.Close()
	io.Copy(Newfile, Oldfile)
}
