package main

import "fmt"

const Pi float32 = 3.1415926

func main() {
	fmt.Printf("%T\n", "Hello world")
	fmt.Printf("%T\n", Pi)
	fmt.Printf("%T\n", 0.111111)
	// 配列
	var array, barray [5]int
	array[0] = 3
	barray = array
	fmt.Printf("before A array = %d \nbefore B array = %d\n", array, barray)
	var i int
	for i < 5 {
		array[i] = i
		i += 1
	}
	//配列は独立しているので、かきかわらない。
	fmt.Printf("after A array = %d \nafter B array = %d\n", array, barray)

	//sliceの値参照
	var aslice, bslice []int
	aslice = array[:3]
	bslice = aslice[:]
	i = 0
	fmt.Printf("before slice = %d\n", aslice)
	fmt.Printf("before bslice = %d\n", bslice)
	for i < 3 {
		aslice[i] += 1
		i += 1
	}
	//sliceで宣言すると、値参照するのでbsliceの値も書き変わる。
	fmt.Printf("array = %d\n", array)
	fmt.Printf("after slice = %d\n", aslice)
	fmt.Printf("after bslice = %d\n", bslice)

	//map
	//keyを文字列で宣言します。値はintとなるディクショナリです。この方法は使用される前にmakeで初期化される必要があります。
	// var b_numbers map[string]int
	//もう一つの宣言方法?????????
	b_numbers := make(map[string]int)
	b_numbers["one"] = 1
	b_numbers["ten"] = 10
	b_numbers["three"] = 3

	fmt.Println("3つ目の数字は：　", b_numbers["three"]) //データの取得
	//"3つ目の数字は: 3"という風に出力されます。
}
