package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Angka", i)
		} else {
			fmt.Println("Continue")
			continue
		}

		if i >= 7 {
			fmt.Println("Break")
			break
		}
	}
}
