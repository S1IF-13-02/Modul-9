package main

import "fmt"

func main() {
	var jumlahOrang int
	fmt.Scan(&jumlahOrang)

	var jumlahMotor int = jumlahOrang / 2

	if jumlahOrang%2 != 0 {
		jumlahMotor = jumlahMotor + 1
	}

	fmt.Println(jumlahMotor)
}
