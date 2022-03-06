package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//fmt.FprintfはCのprintfみたいな感じで書き出すことができる
	//Goでは%vフォーマットを使うとなんでも表示できる。すごすぎぃ！
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
