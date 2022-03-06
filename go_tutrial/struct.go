package main

import "fmt"

type person struct {
	name string
	age  int
}

//2人の年齢を比較します。年齢が大きい方の構造体を返し、年齢の差分も返す
func Older(p1, p2 person) (O_struct person, age_diff int) {
	if p1.age > p2.age {
		age_diff = p1.age - p2.age
		O_struct = p1
		return
	}
	age_diff = p2.age - p1.age
	O_struct = p2
	return
}

func main() {
	var tom person

	//初期値の代入
	tom.name, tom.age = "TOM", 30

	//2つのフィールドを明確にすれば順序は関係なく初期化できる
	bob := person{age: 40, name: "Bob"}

	//順序通りに初期化すれば、フィールド名は省略できる
	paul := person{"Paul", 20}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)

}
