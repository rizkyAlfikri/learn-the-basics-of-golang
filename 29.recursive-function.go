package main

import "fmt"

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(10)
	fmt.Println(recursive)
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
