package main

import "fmt"

func main() {
	someMethods()
}


func someMethods() {

	intArray := add(10)

	fmt.Println(intArray)

	factorialResult := 4

	fmt.Println(factorial(factorialResult))

	fmt.Println(recursiveFactorial(factorialResult))
}

func add(v int) []int {
	intArray := []int{}
	for i := 0; i < v; i++ {
		if i != 0 {
			intArray = append(intArray, i*2)
		}
	}
	return intArray
}

func factorial(v int) int {
	if v < 0 {
		fmt.Printf("v cannot be lower than 0")
		return -1
	}

	if v == 1 || v == 0 {
		return 1
	}

	factorialResult := 1

	for i := 1; i <= v; i++ {
		factorialResult *= i
	}
	return factorialResult
}

func recursiveFactorial(v int) int {
	if v < 0 {
		fmt.Printf("v cannot be lower than 0")
		return -1
	}

	if v == 0 || v == 1 {
		return 1
	}

	return v * recursiveFactorial(v-1)
}
