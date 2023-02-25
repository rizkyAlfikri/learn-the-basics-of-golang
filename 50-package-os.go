package main

import (
	"fmt"
	"os"
)

/*
	Package Os
	Package OS berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
*/

func main() {
	// Mengambil argument
	args := os.Args
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])

	// Mengambil nama sistem operasi yang dipakai
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("hostname", name)
	} else {
		fmt.Println("Error", err)
	}

}
