package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1              // Pass by value
	var address3 *Address = &address1 // pass by reference
	var address4 *Address = &address1
	var address5 *Address = &address1

	address2.City = "Jakarta"
	address3.Province = "Sumatra Barat"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// akan membuat pointer baru, address1 tidak akan berubah
	address4 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// Semua variable yang meng refer ke address1 akan berubah
	*address5 = Address{"Padang", "Sumatra Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)

	var address6 *Address = new(Address)
	fmt.Println("address6", address6) // hasilnya kosong
	address6.City = "Jombang"
	fmt.Println("address6", address6)
}
