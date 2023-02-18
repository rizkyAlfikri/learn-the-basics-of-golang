package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name":    "Budi",
		"address": "Subang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"]) // null

	person["age"] = "20"
	fmt.Println(person["age"])

	// map function
	fmt.Println("Before delete", person)
	fmt.Println("panjang map", len(person))
	delete(person, "age")
	fmt.Println("after delete", person)
	fmt.Println("panjang map", len(person))
}
