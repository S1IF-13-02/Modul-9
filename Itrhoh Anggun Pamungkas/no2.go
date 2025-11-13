package main

import "fmt"

func main() {
	var angka int
	var hasil string = "bukan"

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	if angka < 0 && angka % 2 == 0 {
		hasil = "genap negatif"
	}
	fmt.Println(hasil)
}
