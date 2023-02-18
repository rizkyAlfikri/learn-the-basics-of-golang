package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"Febuary",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"July",
		"Augustus",
		"Septemebr",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println("total data", len(slice1))
	fmt.Println("capacity data", cap(months))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Mei Lago"
	fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Desember baru") // ketika meng append tetapi melebihi batas capacity -> akan membuat array baru -> ketika capacity nya masih sanggung menampung, maka data baru yang di append akan mengreplace element array
	fmt.Println(slice3)

	slice3[2] = "Desember lama"
	fmt.Println(slice3)
	fmt.Println("total data slice2", len(slice2))
	fmt.Println("capacity data slice2", cap(slice2))
	fmt.Println("total data slice3", len(slice3))
	fmt.Println("capacity data slice3", cap(slice3))
	fmt.Println(slice2)
	fmt.Println(months)

	// Membuat array dari awal
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Budi"
	newSlice[1] = "utomo"
	fmt.Println(newSlice)

	/// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	intArray1 := [5]uint8{1, 2, 3, 4, 5}
	intArray2 := [...]uint8{1, 2, 3, 4, 5}
	intSlice := []uint8{1, 2, 3, 4, 5}

	fmt.Println(len(intArray1))
	fmt.Println(cap(intArray2))
	fmt.Println(intSlice)
}
