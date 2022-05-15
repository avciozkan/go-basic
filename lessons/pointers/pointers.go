package main

import (
	"fmt"
	"strings"
)

type String *string

func main() {
	var rstr String
	var pstr string

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "go Türkiye"
	rstr = &pstr
	fmt.Println()
	fmt.Println(rstr)  // show memory address
	fmt.Println(*rstr) // show memory address value pstr
	fmt.Println(pstr)

	fmt.Println()
	replaceStr(&pstr)
	fmt.Println(*rstr)
	fmt.Println(pstr)

	fmt.Println()
	replaceStrCpy(pstr)
	fmt.Println(*rstr)
	fmt.Println(pstr)

	fmt.Println()
	replaceStrCpy2(rstr)
	fmt.Println(*rstr)
	fmt.Println(pstr)
}

func replaceStr(s *string) {
	*s = strings.Replace(*s, "go", "GO", 1)
}

func replaceStrCpy(s string) {
	s = strings.Replace(s, "GO", "111GO", 1)
}

func replaceStrCpy2(s String) {
	*s = strings.Replace(*s, "GO", "11go", 1)
}
