package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Budi"
	names[1] = "Utomo"
	names[2] = "Okta"

	fmt.Println(names[0], names[1], names[2])

	var age = [3]uint8{40, 23, 53}
	fmt.Println(age)
	fmt.Println(age[0], age[1], age[2])

	result := [...]bool{true, false, true}

	fmt.Println(result)
	fmt.Println(len(result))
	fmt.Println(len(names))
	fmt.Println(len(age))

	result[2] = false
	fmt.Println(result)
}
