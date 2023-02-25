package main

import "fmt"

type Filter func(string) string

func main() {
	result := sayWithFilter("Budi", spamFilter)
	badResult := sayWithFilter("Anjing", spamFilter)

	fmt.Println(result)
	fmt.Println(badResult)
}

func sayWithFilter(name string, filter Filter) string {
	return "hello " + filter(name)
}

// func sayWithFilter(name string, filter func(string) string) string {
// 	return "hello " + filter(name)
// }

func spamFilter(name string) string {
	if name == "Anjing" {
		return ". . ."
	} else {
		return name
	}
}
