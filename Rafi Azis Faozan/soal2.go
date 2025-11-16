package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	hasil := "Bukan"
	if x < 0 && x%2 == 0 {
		hasil = "Genap negatif"
	}
	fmt.Print(hasil)
}
