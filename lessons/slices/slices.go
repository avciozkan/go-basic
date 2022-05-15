package main

import "fmt"

var slc []int

func main() {
	slc1 := make([]int, 0, 3)
	slc1 = append(slc1, 1) //slc[0] yapilamaz cunku boyutsuzdur. append fonksiyonu ile atama yapilabilir hale geliyor.

	slc = append(slc, 1, 2, 3, 4, 5, 6, 7, 7777, 7, 7, 7, 77, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7)
	fmt.Println(slc)
	fmt.Println(slc1)
	fmt.Println("Len: ", len(slc), "Cap: ", cap(slc))
	fmt.Println("Len: ", len(slc1), "Cap: ", cap(slc1))
}
