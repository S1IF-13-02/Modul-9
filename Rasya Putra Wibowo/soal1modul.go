package main

import "fmt"

func main() {
	var orang int
	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&orang)

	jumlahMotor := orang / 2
	if orang%2 != 0 {
		jumlahMotor++
	}

	fmt.Println("Jumlah motor yang diperlukan:", jumlahMotor)
}
