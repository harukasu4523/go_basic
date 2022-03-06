package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var test string
	flag.StringVar(&name, "name", "newfile.txt", "you can name newfile")
	flag.StringVar(&test, "test", "test.txt", "test test test")
	flag.Parse()
	newName := flag.Args()
	fmt.Printf("name = %v newname = %v test = %v", name, newName[0], test)

}
