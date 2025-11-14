package main

import "fmt"

func main() {
	var a int
	var hasil = "positif"

	fmt.Print("masukan angka: ")
	fmt.Scan(&a)
	if a <= 0 {
		hasil = "negatif"

	}
	fmt.Print(hasil)
}
