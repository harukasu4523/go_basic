package main

import (
	"net/http"
	"os"
)

func main() {
	//クライアント側のリクエストを送る時とかに使える。
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
