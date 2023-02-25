package main

import "fmt"

func main() {
	// we can assign variable with function

	hello := helloName

	fmt.Println(hello("Budi"))
}

func helloName(name string) string {
	return "hello " + name
}
