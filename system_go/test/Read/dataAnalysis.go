package main

import (
	"fmt"
	"strings"
)

var sources = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(sources)
	var i int
	var f, g float64
	var s string
	// 対応するデータ型に変換して入れてくれる
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v", i, f, g, s)
}
