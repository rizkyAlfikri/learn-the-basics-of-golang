package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
	Package container / ring adalah implementasi structur data circular list
	Circular list adalah data ring, dimana diakhir element akan kembali ke element awal (HEAD)
*/

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next() // move to next ring
	}

	// tidak bisa looping dengan loop, harus menggunakan method bawaan
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
