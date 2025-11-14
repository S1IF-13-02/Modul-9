package main

import "fmt"

func main() {
	var jumlahOrang, jumlahMotor int

	fmt.Print("jumlah orang yang touring: ")
	fmt.Scan(&jumlahOrang)
	
	jumlahMotor = jumlahOrang / 2

	if jumlahOrang%2 != 0 {
		jumlahMotor++
	}

	fmt.Println("Jumlah motor :", jumlahMotor)
}