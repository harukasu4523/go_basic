package main

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"os"
)

func readChunks(file *os.File) []io.Reader {
	// chunkを格納する配列
	var chunks []io.Reader

	// 最初の8バイトを飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length+12)))
		// 次のチャンクのスタート位置に移動
		// 現在地は長さを読み終わった後なので
		// チャンク名(4byte) + データ長(length) + CRC(4byte)先に移動
		// 第二引数に1を入れることで前回のseekからの移動になる。
		// 0 = 先頭から: 1 = 今のSeek位置から: 2 = 末尾から(この場合マイナスじゃないと動かない)
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func textChunk(text string) io.Reader {
	//tExtというチャンクを作るメソッド
	byteData := []byte(text)
	var buffer bytes.Buffer
	// ASCIIデータの長さををBigEndianでbufferに書き込む
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	// チャンク名の書きこみ
	buffer.WriteString("tEXt")
	// 実データの書き込み
	buffer.Write(byteData)
	// CRCを計算して追加
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	crc.Write(byteData)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer

}

func main() {
	file, err := os.Open("Lenna_(test_image).png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	chunks := readChunks(file)
	// シグニチャに書き込み
	io.WriteString(newFile, "\x39PNG\r\n\x1a\n")
	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, chunks[0])
	// テキストチャンクを追加
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}
