package main

import "fmt"

// Varags in kotlin

func main() {
	result := sumAll(5, 10, 10, 10, 10, 10)
	fmt.Println("num params", result)

	// slice
	numbersSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice params", sumAll(4, numbersSlice...))

	// array
	numberArray := [...]int{2, 4, 6, 8, 10}
	fmt.Println("array params", sumAll(1, numberArray[:]...))
}

func sumAll(initNum int, numbers ...int) int {
	total := initNum
	for _, number := range numbers {
		total += number
	}

	return total
}
