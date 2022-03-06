package main

import (
	"bufio"
	"os"
)

func main() {
	//bufio構造体はFlushメソッドをつかうまで文字列をため込める
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
