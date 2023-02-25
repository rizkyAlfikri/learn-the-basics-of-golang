package main

import "fmt"

func main() {
	customer := Customer{
		Name:    "Budi",
		Address: "Bandung",
		Age:     28,
	}

	customer2 := Customer{"Budi", "Bandung", 28, false}
	fmt.Println(customer2)

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	var shinta Customer
	shinta.Name = "Shinta"
	shinta.Address = "Indonesia"
	shinta.Age = 22

	fmt.Println(shinta)
}

type Customer struct {
	Name, Address string
	Age           uint8
	Maried        bool
}
