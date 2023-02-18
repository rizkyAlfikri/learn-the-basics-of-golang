package main

import "fmt"

func main() {
	sayHelloTo("Budi", "Widodo")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
