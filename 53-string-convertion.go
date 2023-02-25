package main

import (
	"fmt"
	"strconv"
)

func main() {
	booleanResult, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(booleanResult)
	} else {
		fmt.Println(err.Error())
	}

	// number that will converted,
	// baseSize = Desimal, biner, okta, hexa
	// Number size = int8, int16, int32, int64
	numberResult, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(numberResult)
	} else {
		fmt.Println(err)
	}

	value := strconv.FormatInt(1000000, 16)
	fmt.Println(value)

	valueInt, err := strconv.Atoi("1000000")
	fmt.Println(valueInt)
}
