package main

import "fmt"

var array0 [3]int

var array1 = [5]int{1, 2, 3, 4, 5}

var arrayString [5]string

func main() {
	array2 := make([]int, 19)
	array2[0] = 1
	fmt.Println(array0)
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(arrayString)
	fmt.Print("len : ", len(arrayString))
}
