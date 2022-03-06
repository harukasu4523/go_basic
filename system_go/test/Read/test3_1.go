package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// var argc int
	var name string
	// 自作オプションの作成flag.データ型(結果をいれる変数, オプション, 初期値, 説明)
	flag.StringVar(&name, "n", "new.txt", "新しく作るファイルの名前を指定する")
	// 適用する
	flag.Parse()

	var oldFile string
	//　オプションでずれ込んだ分とそうでない時で場合わけする
	if os.Args[1] == "-n" {
		oldFile = flag.Args()[0]
		//オプション含めて引数が3個以外の時にエラー
		if len(os.Args[1:]) != 3 {
			fmt.Println("Incorrect number of arguments")
			os.Exit(1)
		}
	} else {
		oldFile = os.Args[1]
	}

	file, err := os.Open(oldFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// nameオプションで与えたれたファイルをつくる
	newFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	// コピーする
	io.Copy(newFile, file)
}
