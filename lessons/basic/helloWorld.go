package main

import (
	"fmt"
)

const sayi6 = 3
const (
	a = 1
	b = 2
	c = 3
)

const (
	d = iota
	e = 2
	f = 3
)

var sayiGlobal = 4

func main() {
	var sayi1 = 1
	var sayi2 = 3
	sayi3 := 6
	fmt.Println("hello world!")
	fmt.Println(add(sayi1, sayi2))
	fmt.Println(add(sayi1, sayi3))
	fmt.Println(add(sayi1, sayiGlobal))
	fmt.Println(add(add(a, b), c))
	fmt.Println(add(add(d, e), f))
}

func add(a int, b int) int {
	return a + b
}
