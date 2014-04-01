package main

import (
	"fmt"
)

type Head struct {
	Eyes  string
	Nose  string
	Mouth string
	Ear   string
}

type Pussy struct {
	Color string
	Size  int
	Shape string
}

type Body struct {
	Head   Head
	Breast string
	Hands  string
	Ashole string
	Pussy  Pussy
	Feet   string
}

type Women struct {
	Name string
	Age  int
	Body Body
}

func (w *Women) MakeLove() {
	fmt.Println("I am ", w.Age, " now and my name is", w.Name)
	fmt.Println("er...")
}

func main() {
	women := Women{Name: "Nancy", Age: 18}
	women.MakeLove()
}
