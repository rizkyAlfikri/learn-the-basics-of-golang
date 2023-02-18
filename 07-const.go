package main

import "fmt"

func main() {
	const firstName string = "Budi"
	const lastName = "Otomo"
	const age = 65
	const name string = firstName + " " + lastName

	fmt.Println(name+" ", age)

	const (
		firstName1       = "dodo"
		lastName1  uint8 = 4
	)

	fmt.Println(firstName, " ", lastName1)
}
