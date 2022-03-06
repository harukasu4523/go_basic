package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	// インターフェースの型をもつ変数を宣言
	var talker Talker
	var men = Greeter{"Jon"}
	// インターフェースを満たす構造体のポインタは代入できる
	talker = &Greeter{"wozozo"}
	talker.Talk()
	men.Talk()
}
