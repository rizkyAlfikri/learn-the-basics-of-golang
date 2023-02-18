package main

import "fmt"

func main() {
	name := "Budi"

	if name == "Budi" {
		fmt.Println("Hello", name)
	} else if name == "utomo" {
		fmt.Println("My name is", name)
	} else {
		fmt.Println("Hello who are you ?")
	}

	// If short statement
	if length := len(name); length >= 5 {
		fmt.Println("Lebih sama dengan dari 5 ", length)
	} else {
		fmt.Println("Kurang dari 5", length)
	}
}
