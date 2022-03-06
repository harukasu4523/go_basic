package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Header().Set()でサーバーで受け取れるデータ規約を決めれるっぽい
	// 今回の場合だとJSONをgzip化したものを表現できる
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化するもとのデータ
	source := map[string]string{
		"Hello": "World",
	}
	// 以下コード記入
	// 書き込むfd(ファイル識別子)に対してgzip化できる構造体を返している?
	gw := gzip.NewWriter(w)
	gw.Header.Name = "test.txt"
	// 構造体とかをjsonにしてくれるメソッド
	outputjson, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}
	// 異なるfdに対して同じことを書き込む
	writer := io.MultiWriter(gw, os.Stdout)
	// 書き込む
	i, err := io.WriteString(writer, string(outputjson))
	if err != nil {
		panic(err)
		fmt.Fprintf(os.Stdout, "%v", i)
	}
	// httpの方なのかgzipのほうなのかはわからないけどどっちかがbuffer持ってるっぽいので
	// Flushで吐き出してあげる
	gw.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
