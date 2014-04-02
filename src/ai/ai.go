package main

import (
	"fmt"
)

func test(s string) {
	fmt.Println(s)
}

func main() {
	defer test("defer 1")
	defer test("defer 2")
	defer test("defer 3")
	fmt.Println("ai")
}
