package main

import "fmt"

// Recover adalah function yang digunakan untuk menangkap data panic
// Dengan recover proses panic akan berhenti, sehingga program akan tetap berjalan

/**
sd
asd
*/
func main() {
	runApp(true)
	fmt.Println("runApp has finished")
}

func runApp(isError bool) {
	defer endApp()

	if isError {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi berjalan")
}

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("Aplikasi selesai")
}
