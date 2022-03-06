package main

import (
	"fmt"
)

type testInt func(int) bool //関数の型宣言

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// ここでは宣言する関数の型を一つの引数とみなします。

func filter(slice []int, f testInt) []int {
	var result []int
	b_slice := slice[2:4]
	//range文で0からのindex順に配列ないの値を取り出して、valueに加える
	//今回はindexを特には使わないので_で消しておく
	for _, value := range b_slice {
		if f(value) {
			//append(結果を返すスライス, 加える数字なりなんなり)第一引数に第二引数を加えて返す関数
			result = append(result, value)
		}
		fmt.Println("b_slice is : ", value)
	}

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}
