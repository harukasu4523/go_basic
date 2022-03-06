package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// Writeメソッドで書き込まれた内容をためておき後で結果を受け取れるbytes.Bufferがある
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example"))
	fmt.Println(buffer.String())
	// WriteStringはio.Writerのメソッドではないので他の構造体では使えません
	//インターフェースのリピーターに使われていない????ッてこと?
	buffer.WriteString(" bytes.Buffer example")
	io.WriteString(&buffer, " bytes.Buffer example")
	fmt.Println(buffer.String())

}
