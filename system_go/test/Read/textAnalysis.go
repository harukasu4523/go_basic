package main

import (
	"bufio"
	"fmt"
	"strings"
)

// バッククォートで書式をそのまま保存できる
var source = `1行目
2行目
3行目 `

// func main() {
// 	// string.Reader型の構造体をそのままbufio.Reader型構造体にして返す
// 	reader := bufio.NewReader(strings.NewReader(source))
// 	for {
// 		line, err := reader.ReadString('\n')
// 		fmt.Printf("%v\n", line)
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }
func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
