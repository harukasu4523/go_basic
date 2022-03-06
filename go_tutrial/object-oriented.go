package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, hight float64
}

type Circle struct {
	redius float64
}

//func + 構造体宣言でそのstructのフィールドを利用することができる
//
func (r Rectangle) area() float64 {
	return r.width * r.hight
}

func (c Circle) area() float64 {
	return c.redius * c.redius * math.Pi
}

func main() {

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	//r1・r2のareaはrectangleのareaになる
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	//c1・c2のareaはCircleのareaになる
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

}
