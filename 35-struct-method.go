package main

import "fmt"

func main() {
	customer := Customer{
		Name:    "Budi",
		Address: "Bandung",
		Age:     28,
	}

	customer.sayHi("Eko")
}

type Customer struct {
	Name, Address string
	Age           uint8
	Maried        bool
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
