package main

import "fmt"

func main() {
	counter := 1
	var intArray [10]int

	for counter <= 10 {
		fmt.Println("Counter yang terhitung", counter)
		if counter <= 6 {
			intArray[counter-1] = counter
		}
		counter++
	}

	fmt.Println(intArray)
	fmt.Println(len(intArray))
	fmt.Println(cap(intArray))

	// for
	for counter2 := 0; counter2 < 10; counter2++ {
		fmt.Println("Counter 2 yang terhitung", counter2)
	}

	// for range
	slice := []string{"Budi", "Utomo", "Widodo"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println("index =", index, "value =", value)
	}

	for _, value := range slice {
		fmt.Println("value =", value)
	}
}
