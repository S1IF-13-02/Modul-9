package main

import "fmt"

func main() {
	var orang int

	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&orang)

	motor := orang / 2

	if orang%2 != 0 {
		motor++
	}

	fmt.Println("Motor yang dibutuhkan:", motor)
}
