package main

import "fmt"

func main() {
	var budi Person
	budi.Name = "budi"
	fmt.Println("Hello", budi.GetName())
	sayHello(budi)

	cat := Animal{
		name: "Tom",
	}

	sayHello(cat)
}

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	name string
}

func (animal Animal) GetName() string {
	return animal.name
}
