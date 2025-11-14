package main

import "fmt"

func main() {
	var a int
	var hasil int

	fmt.Print("masukan jumlah orang yang akan mengikuti touring : ")
	fmt.Scan(&a)

	if a%2 == 0 {
		hasil = a / 2

	}

	hasil = a/2 + a%2
	fmt.Print("total motor yang diperlukan untuk touring adalah : ", hasil)
}
