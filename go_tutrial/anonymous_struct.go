package main

import (
	"fmt"
)

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

// type Test struct {
// 	name string
// }

//匿名フィールドは構造体の中に構造体いれるみたいなやつ
type Student struct {
	Human      //匿名フィールド
	Skills     //匿名フィールド、自分で定義した型 string slice
	int        //ビルトイン型を匿名フィールドとします<-は？？
	speciality string
	// Test       //2つ目の匿名フィールド
}

func main() {
	//学生Janeを初期化します
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	//ここで対応するフィールドにアクセスする
	//匿名フィールドへのアクセスはStudent.Human.nameみたいな形ではなくStudent.nameで直接アクセスできる。
	//もしStudentにnameフィールドがある時は外側つまりStudent側が優先される
	//匿名フィールド内に同じフィールドがある場合はエラー出る(どこ参照したらいいかわからんからね)
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 彼女のskill技能フィールドを修正します
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	//匿名ビルトイン型のフィールドを修正します
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)

}
