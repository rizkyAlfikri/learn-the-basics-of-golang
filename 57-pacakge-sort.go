package main

import (
	"fmt"
	"sort"
)

/*
Package sort adalah pacakage yang berisikan utilitas untuk proses pengurutan
Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort. Interface
*/
func main() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 26},
	}

	fmt.Println("Before sort", users)

	sort.Sort(UserSlice(users))

	fmt.Println("After sort", users)
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

// func (value UserSlice) Sort()
