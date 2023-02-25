package main

import "fmt"

type BlockList func(string) bool

func registerUser(name string, blocklist BlockList) {
	if blocklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}

}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	register := func(name string) {
		registerUser(name, blacklist)
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)
	register("Root")
}
