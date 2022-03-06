package main

import (
	"fmt"
)

type Say interface {
	Greeter()
}

type Human struct {
	name   string
	age    int
	weight float32
}

func (h Human) Greeter() {
	fmt.Printf("my name is %s. I am %d. My weight is %f", h.name, h.age, h.weight)
}

func main() {
	harukasu := Human{"Harukasu", 27, 63.25}
	harukasu.Greeter()
}
