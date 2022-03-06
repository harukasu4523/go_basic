package main

import "fmt"

//inheritance.goの機能を拡張していく
//メソッドの組み合わせがinterfaceと呼ばれている

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名フィールド
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名フィールド
	company string
	money   float32
}

//Human上でメソッドを定義
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call on %s\n", h.name, h.phone)
}

//HumanオブジェクトにSingメソッドを追加
func (h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la lala...", lyrics)
}

//HumanメソッドにGuzzleメソッドを実装
// func (h Human) Guzzle(beerStein string) {
// 	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
// }

//Employeeのフィールドにもアクセスしたい場合はアクセスしたいフィールドでメソッドを定義する
//EmployeeはHumanのSayHiメソッドをオーバーロードします
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

//StudentはBorrowMoneyメソッドを実装します
// func (s Student) BorrowMoney(amount float32) {
// 	s.loan += amount //(again and again and...)
// }

//interfaceを定義
//Interface Men はHuman, Student, Employeeに実装される
//3つの型はSayHiとSingを実装するから。
type Men interface {
	SayHi()
	Sing(lyrics string)
	// Guzzle(beerStein string)
}

// type YoungChap interface{
// 	SayHi()
// 	Sing(song string)
// 	BorrowMoney(amount float32)
// }

// type ElderlyGent interface{
// 	SayHi()
// 	Sing(song string)
// 	SpendSalary(amount float32)
// }

func main() {
	mike := Student{Human{"Mike", 18, "222-555-555"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 24, "333-111-111"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 45, "555-999-888"}, "golang.inc", 1000}
	tom := Employee{Human{"Tom", 44, "333-444-555"}, "Things Ltd", 5000}

	//Men型の変数iを定義
	var i Men

	//iにはStudentを保存できる
	i = mike
	fmt.Println("This is Mike, a Student")
	i.SayHi()
	i.Sing("Born to be wild")
	fmt.Printf("my name is%s\n", mike.name)

	//iはEmployeeも保存できる。Menがそれぞれに実装されているからね
	i = tom
	fmt.Println("This is Tom, an Employee")
	i.SayHi()
	i.Sing("November rain")

	//sliceのMenを定義します
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//この3つはどれも異なる要素ですが同じインターフェースを実装しています
	///インターフェース内のメソッドに対応する構造体を代入してあげることでメソッドを利用することができる。
	x[0], x[1], x[2] = paul, sam, tom

	fmt.Println("hello my name is ", x[0])
	for _, value := range x {
		value.SayHi()
	}

}
