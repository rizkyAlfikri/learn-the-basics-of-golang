package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Budi Utomo", "Budi"))
	fmt.Println(strings.Contains("Budi Utomo", "Joko"))
	fmt.Println(strings.Split("Budi Utomo Setiawan", " "))

	fmt.Println(strings.ToLower("Budi Utomo"))
	fmt.Println(strings.ToUpper("Budi Utomo"))

	fmt.Println(strings.ToTitle("budi utomo setiawan"))

	fmt.Println(strings.Trim("    Budi Utomo Setiawan   ", " "))
	fmt.Println(strings.Trim("aaaaaaBudi Utomo Setiawan aaaaaa", "a"))

	fmt.Println(strings.ReplaceAll("Budi Budi Budi", "Budi", "Joko"))

}
