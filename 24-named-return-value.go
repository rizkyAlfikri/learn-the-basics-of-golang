package main

import "fmt"

func main() {
	fmt.Println(getCompleteName())
	fmt.Println(getCompleteNameVer2()) // -> Result will be same
}

func getCompleteName() (firstName, lastName string, age uint8) {
	firstName = "Budi"
	lastName = "Widodo"
	age = 42

	return firstName, lastName, age
}

func getCompleteNameVer2() (firstName, lastName string, age uint8) {
	firstName = "Budi"
	lastName = "Widodo"
	age = 42

	return
}
