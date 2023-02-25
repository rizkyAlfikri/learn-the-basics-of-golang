package main

import (
	"fmt"
)

// type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// fitur ini sering digunakan ketika kita bertemu dengan data interface kosong

func main() {
	// result := random()

	// resultStr := result.(string)
	// fmt.Println(resultStr)

	result2 := random()
	switch value := result2.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

func random() interface{} {
	return 123
}
