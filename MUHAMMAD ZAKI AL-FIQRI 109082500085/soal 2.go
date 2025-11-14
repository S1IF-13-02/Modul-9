package main

import "fmt"

func main() {
	var a int
	hasil := "bukan genap negatif"

	fmt.Print("masukan angka : ")
	fmt.Scan(&a)

	if a%2 == 0 && a < 0 {
		hasil = "genap negatif"
	}
	fmt.Print(hasil)
}
