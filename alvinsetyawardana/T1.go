package main

import "fmt"

func main() {
	var jumlahorang int
	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&jumlahorang)
	jumlahMotor := (jumlahorang + 1) / 2
	fmt.Printf("Jumlah motor yang dibutuhkan: %d\n", jumlahMotor)

}
