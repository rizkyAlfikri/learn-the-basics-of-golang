package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Utomo":
		fmt.Println("hello Utomo")
	case "Budi":
		fmt.Println("Hello Budi")
	case "Eko":
		fmt.Println("Hello eko")
		fmt.Println("Hello eko2")
	default:
		fmt.Println("Who are you")
	}

	// short statement

	switch length := len(name); length {
	case 5:
		fmt.Println("Sama dengan 5")
	case 6:
		fmt.Println("sama dengan 6")
	default:
		fmt.Println("panjang nama adalah", length)
	}

	// switch tampa expression
	length2 := len(name)
	switch {
	case length2 > 5:
		fmt.Println("Lebih dari 5")
	case length2 < 5:
		fmt.Println("Kurang dari 5")
	default:
		fmt.Println("sama dengan 5")
	}

}
