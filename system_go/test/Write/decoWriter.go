package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	// file2, err := os.Create("multiwriter2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//MultiWriterは引数にファイル情報を与えるとそれぞれに書き出す
	//引数は可変長でfile以上
	writer := io.MultiWriter(file)
	io.WriteString(writer, "io.MultiWriter example")
}
