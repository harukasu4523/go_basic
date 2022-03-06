package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	// 指定されたネットワーク上のアドレスに接続するnet.Dial(プロトコル、アドレス)
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	//リクエストをサーバーに投げて、io.Reader使ってレスポンスを受け取っている
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダーを表示してみる。
	fmt.Println(res.Header)
	//ボディーを表示してみる。最後にはClose()すること
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
