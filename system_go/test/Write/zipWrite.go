package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	//gzip拡張子のファイルを作成
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	//作ったファイルにたいしてgzip処理をするメソッド
	writer := gzip.NewWriter(file)
	//ジップ処理をするファイルの名前の変更
	writer.Header.Name = "test.txt"
	//ファイル内に書き込む
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}
