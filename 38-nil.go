package main

import "fmt"

// nil == null

func main() {
	person := NewMap("Budi")
	fmt.Println(person)

	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
