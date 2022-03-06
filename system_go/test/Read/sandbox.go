package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("args count is: %d\n", len(os.Args))
	fmt.Printf("args : %v", os.Args)
	for i, v := range os.Args {
		fmt.Printf("Args[%d]: %v\n", i, v)
	}
}
