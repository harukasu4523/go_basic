package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名フィールド
	school string
}

type Employee struct {
	Human   //匿名フィールド
	company string
}

//Human上でメソッドを定義
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call on %s\n", h.name, h.phone)
}

//Employeeのフィールドにもアクセスしたい場合はアクセスしたいフィールドでメソッドを定義する
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 18, "222-555-555"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "555-999-888"}, "golang.inc"}

	mark.SayHi()
	sam.SayHi()
}
