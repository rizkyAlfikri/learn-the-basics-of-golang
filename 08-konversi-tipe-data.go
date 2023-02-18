package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println("nilai32 ", nilai32)
	fmt.Println("nilai64 ", nilai64)
	fmt.Println("nilai8 ", nilai8)

	var name = "Budi"
	var firstChar uint8 = name[0]
	var firstCharString = string(firstChar)
	fmt.Println(firstCharString)
}
