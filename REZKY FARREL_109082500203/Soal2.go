package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	if n < 0 && n % 2 == 0 {
		fmt.Println("Genap Negatif")
	} else {
		fmt.Print("Bukan")
	}
}