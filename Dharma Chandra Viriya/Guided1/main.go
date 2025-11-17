package main

import "fmt"

func main() {
	var x int64

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&x)

	if x < 0 {
		x = -x
	}

	fmt.Println(x)
}
