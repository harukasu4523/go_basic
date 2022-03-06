package main

import "fmt"

func main() {
	//aを空のインターフェースとして定義
	var a interface{}
	var i int = 5
	s := "Hello World"
	// aは任意な型の数値を保存できます
	//C言語のvoidポインタ型みたいなやつ
	//int型
	a = i
	fmt.Println(a)
	//string型
	a = s
	fmt.Println(a)

}
