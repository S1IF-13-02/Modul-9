package main

import "fmt"

func main() {
	var jumlahOrang int
	fmt.Print("Masukkan jumlah orang yang akan touring: ")
	fmt.Scan(&jumlahOrang)

	jumlahMotor := jumlahOrang / 2
	if jumlahOrang%2 != 0 {
		jumlahMotor++
	}

	fmt.Println("Jumlah motor yang diperlukan:", jumlahMotor)
}
