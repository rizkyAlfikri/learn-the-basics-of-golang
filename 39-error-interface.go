package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println("result", result)
	} else {
		fmt.Println("error", err)
	}
}

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}
