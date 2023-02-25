package main

import "fmt"

// Defer adalah sebuah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

func main() {
	runApplication(0)
}

func runApplication(num int) {
	// call defer at the first line in scope method
	defer logging() // Will be called after runApplication finish, even runAplication throw exception
	fmt.Println("Run Application")
	result := 10 / num
	fmt.Println("Value", result)
}

func logging() {
	fmt.Println("Selesai memanghil function")
}
