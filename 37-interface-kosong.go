package main

import "fmt"

// interface kosong sama dengan any di kotlin

func main() {
	result := Ups(1)
	fmt.Println(result)
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}
