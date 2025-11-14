package main

import (
	"fmt"
	"math"
)	
func main() {
	var jumlahOrang float64
	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&jumlahOrang)

	if jumlahOrang > 0 {
		jumlahOrang /= 2
		motor := math.Round(jumlahOrang)
		fmt.Print("jumlah motor: ", motor)
	} else {
		fmt.Println("Tidak Sesuai")
	}
}