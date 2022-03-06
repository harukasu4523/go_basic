package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	c := 1024
	randam := make([]byte, c)
	_, err := rand.Read(randam)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("randam.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, string(randam))
}
