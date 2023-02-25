package main

import "fmt"

// Panic function adalah function yang kita gunakan untuk menghentikan program
// panic function biasanya dipanggil ketika terjadi error pada saat program berjalan
// Saat panic function dipanggil, program akan berhenti, namun defer function akan tetap dieksekusi

func main() {
	runApp(true)
}

func runApp(isError bool) {
	defer endApp()

	if isError {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi berjalan")
}

func endApp() {
	fmt.Println("Aplikasi selesai")
}
