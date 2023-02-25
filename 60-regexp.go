package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("Budi"))
	fmt.Println(regex.MatchString("bumi"))
	fmt.Println(regex.MatchString("Budhaphest"))

	search := regex.FindAllString("bumi budi budha bubi", 1)
	fmt.Println(search)
}
