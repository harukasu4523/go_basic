package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// 対象のbinary file(今回だとpng)のチャンクの長さをとりだす

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	// 第二引数と等しい時
	if bytes.Equal(buffer, []byte("tEXt")) {
		// byteスライスの変数chunkの中身がbyte型だからそれを用意する
		rawText := make([]byte, length)
		// Readでchunkの中身をrawTextに入れる
		chunk.Read(rawText)
		// stringにキャストして出力
		fmt.Println(string(rawText))
	}
}
func readChunks(file *os.File) []io.Reader {
	//チャンクを格納する配列
	var chunks []io.Reader

	//最初の8バイトを飛ばす
	file.Seek(8, 0)
	var offset int64 = 8
	for {
		var length int32
		// チャンクの先頭4byteが保持しているデータの長さをもっているので
		//binary.Readでシグニチャを抜いた(上のfile.Seek)長さデータの4byteの読み込みをする
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		// io.NewSectionReaderではファイルの先頭ポインタから始まるので、
		// 一回目はシグニチャを飛ばすオフセット8byteめからスタートしている
		// file.Seekの返り値がおそらく移動したbyte数なので次のチャンクのスタート位置になってるってわけ
		chunks = append(chunks,
			io.NewSectionReader(file, offset, int64(length)+12))
		// 次のチャンクの先頭に移動
		// 現在の位置は長さを読み終わった箇所なので
		// チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
