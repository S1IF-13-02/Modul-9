package main

import "fmt"

func main() {
	var hasil int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&hasil)
	if hasil < 0 {
		hasil = -hasil
	}
	fmt.Print(hasil)
}
