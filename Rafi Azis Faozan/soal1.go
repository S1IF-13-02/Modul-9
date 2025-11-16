package main

import "fmt"

func main() {
	var orang int
	fmt.Print("Masukkan orang: ")
	fmt.Scan(&orang)
	motor := orang / 2
	if orang%2 == 1 {
		motor = motor + 1
	}
	fmt.Print("Jumlah motor yang dibutuhkan: ", motor)
}
