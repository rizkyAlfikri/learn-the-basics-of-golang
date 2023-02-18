package main

import "fmt"

func main() {
	fmt.Println(getHello("Budi"))
}

func getHello(name string) string {
	var result string
	if name == "" {
		result = "Silahkan masukan nama yang valid"

	} else {
		result = "Hello " + name
	}

	return result
}
