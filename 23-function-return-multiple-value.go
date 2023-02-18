package main

import "fmt"

func main() {
	firstName, lastName, age, _ := getFullName()
	fmt.Println(firstName, lastName, age)

}

func getFullName() (string, string, uint8, bool) {
	return "Budi", "Widodo", 29, false
}
