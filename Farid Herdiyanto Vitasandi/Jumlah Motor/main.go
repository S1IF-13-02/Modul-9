package main

import "fmt"

func main(){

	var jumlah_orang int

	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&jumlah_orang)
	
	jumlah_motor := jumlah_orang / 2

	if jumlah_orang % 2 != 0 {
		jumlah_motor += 1
	}
	fmt.Print("Jumlah motor yang dibutuhkan adalah: ", jumlah_motor)

}
