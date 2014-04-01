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

type Body struct {
	Head   Head
	Breast string
	Hands  string
	Feet   string
}

type Women struct {
	Name string
	Age  int
	Body Body
}
