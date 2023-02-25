package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address2 := Address{
		City:     "Cianjur",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	ChangeWithPointer(&address1)
	ChangeWithoutPointer(address2)

	fmt.Println("address1", address1) // City berubah
	fmt.Println("address2", address2) // City tidak berubah
}

func ChangeWithPointer(address *Address) {
	address.City = "Malang"
}

func ChangeWithoutPointer(address Address) {
	address.City = "Padang"
}
