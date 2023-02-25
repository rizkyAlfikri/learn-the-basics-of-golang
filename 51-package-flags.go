package main

import (
	"flag"
	"fmt"
)

/*
	Package flag berisikan fungsionalotas untuk memparasing command line argument
*/

func main() {
	var host string
	flag.StringVar(&host, "host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *int = flag.Int("password", 0, "Put your database user")
	var age *int = flag.Int("age", 0, "Put your database user")

	// Mandatory
	flag.Parse()

	fmt.Println("Host", host)
	fmt.Println("User", *user)
	fmt.Println("Password", *password)
	fmt.Println("age", *age)
}
