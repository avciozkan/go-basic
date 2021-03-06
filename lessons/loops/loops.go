package main

import (
	"fmt"
	"time"
)

func main() {
	slc_1 := []int{1, 2, 3}
	slc_2 := []int{}
	slc_3 := []int{}
	//foreach loop
	for i, value := range slc_1 {
		fmt.Printf("index: %d value: %d \n", i, value)
	}

	//classic for loop
	for i := 0; i < 10; i++ {
		slc_2 = append(slc_2, i)
	}

	for i := range [10]int{} {
		slc_3 = append(slc_3, i)
	}

	fmt.Println(slc_2)
	fmt.Println(slc_3)

	for _, value_1 := range slc_2 {
		for _, value_2 := range slc_1 {
			value_1 += value_2
		}
	}

	//there is no diff if you use above loops, if you want to change, see the below code section.

	for _, value2 := range slc_2 {
		for i := range slc_1 {
			slc_1[i] += value2
		}
	}
	fmt.Println()
	fmt.Println(slc_1)
	fmt.Println(slc_2)

	// optional (if you want to see the below function what it happened to work, you can remove comment section)
	someMethods()

	// same while loop
	c := time.After(5 * time.Second)
	for {
		b := false
		select {
		case <-c:
			b = true

		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
		if b {
			break
		}
	}
}
