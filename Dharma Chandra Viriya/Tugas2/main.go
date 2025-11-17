package main

import "fmt"

func main() {
	var x int64

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&x)

	if x < 0 && x%2 == 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}
