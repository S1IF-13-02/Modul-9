package main

import "fmt"

func main() {
	var angka int
	var isPositif bool = true

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	if angka < 0 {
		isPositif = false
	}
	fmt.Print("Apakah angka tersebut positif? ", isPositif)
}
	