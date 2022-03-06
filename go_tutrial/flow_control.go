package main

import (
	"fmt"
)

//修飾詞 + 変数名 + 仮引数 + 戻り値のデータ型
//戻り値は複数返すことができる
func max(x, y int) int {
	if x < y {
		return (y)
	}
	return (x)
}

//複数の戻り値を取ることもできる。
func max_min(x, y int) (int, int) {
	if x < y {
		return y, x
	}
	return x, y
}

//戻り値に変数名をつけることで、それにいれてreturnするだけで良い
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func main() {
	if x := 15; x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
	// fmt.Println("x is no %d", x)

	//goto文
	// i := 1
	//gotoのタグは大文字小文字を区別する。
	// Here:
	// 	fmt.Println(i)
	// 	i++
	// 	if i == 10 {
	// 		return
	// 	}
	// 	goto Here

	//switch文
	integer := 3
	switch integer {
	case 5:
		fmt.Println("first")
		//fallthroughをつけると当てはまるところ全て表示してくれる。
		fallthrough
	case 3:
		fmt.Println(integer)
		fallthrough
	case 4:
		fmt.Println("aaaaaa")
		fallthrough
	default:
		fmt.Println("default")
	}
	fmt.Printf("%T\n", integer)

	//関数
	var x int = 3
	var y int = 19
	var ans int = 0

	ans = max(x, y)
	fmt.Printf("%d\n", ans)
	//戻り値がふたつなので変数も二つ用意する。
	max, min := max_min(x, y)
	//片方の戻り値がいらない時は”_”で潰せる
	_, test := max_min(x, y)
	fmt.Printf("max = %d min = %d\n", max, min)
	fmt.Printf("%d\n", test)

	sumans, multians := SumAndProduct(x, y)
	fmt.Printf("add = %d, multiplied = %d", sumans, multians)

	//defer文
	//file closeとか関数の終了時に実行してくれる文
	//↓のこんな感じでreturn 後になんかやってくれるerror処理とかに使うといいかもしれん
	// func ReadWrite() bool {
	// 	file.Open("file")
	// 	defer file.Close()
	// 	if failureX {
	// 		return false
	// 	}
	// 	if failureY {
	// 		return false
	// 	}
	// 	return true
	// }
}
