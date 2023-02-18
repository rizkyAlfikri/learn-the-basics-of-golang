package main

import "fmt"

func main() {
	var asd string = "asdasd"
	fmt.Println(asd)

	asd = "1234"
	fmt.Println(asd)

	var number uint8
	number = 255
	fmt.Println(number)

	var dasd string = "asdasd"
	var laldf = 123213
	fmt.Println(dasd)
	fmt.Println(laldf)

	name := "Name"
	num := 31
	fmt.Println(name)
	fmt.Sprintln(num)

	name = "name 2"
	fmt.Println(name)

	var (
		firstName       = "Budi"
		lastName        = "Utomo"
		age       uint8 = 23
	)

	fmt.Println(firstName, " ", lastName, " age = ", age)

	firstName = "Sasha"
	fmt.Println(firstName, " ", lastName, " age = ", age)

}
