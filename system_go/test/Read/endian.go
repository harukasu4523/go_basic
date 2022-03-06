package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32ビットのビッグエンディアンのデータ(10000)
	// int型はbig endianだと0x0, 0x0, 0x27, 0x10と入っていて。
	// little endianに直すと0x10, 0x27, 0x0, 0x0のようになる
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var big, little int32
	//　エンディアンの変換
	// やってることは第3引数のデータ型で条件分岐させて、そのデータ型でキャストする今回は(int32)
	// その後第二引数に指定した構造体がインターフェース型にキャストされ、インターフェース内メソッドを使い
	// byteスライスを一つずつint32でキャストして、下記ビットを左シフトしてor演算するとbigendianになる
	// return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	// big -> little のときはスライスを逆順にする(b[3]から始める)
	binary.Read(bytes.NewReader(data), binary.BigEndian, &big)
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &little)

	fmt.Printf("data of big endian: %d\n", big)
	fmt.Printf("data of little endian: %d\n", little)

}
