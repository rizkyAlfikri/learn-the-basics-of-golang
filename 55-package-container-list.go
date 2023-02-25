package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Budi")
	data.PushBack("Utomo")
	data.PushBack("Setiawan")
	data.PushFront("Shinta")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
