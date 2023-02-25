package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) MarriedWithPointer() {
	man.Name = "Mr. " + man.Name
}

func (man Man) MarriedWithoutPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	budi := Man{"Budi"}

	fmt.Println("Before", budi.Name)
	budi.MarriedWithoutPointer() // Tidak akan mengubah field value
	fmt.Println("After", budi.Name)

	joko := Man{"Joko"}
	fmt.Println("Before", joko.Name)
	joko.MarriedWithPointer() // akan mengubah field value
	fmt.Println("After", joko.Name)

}
