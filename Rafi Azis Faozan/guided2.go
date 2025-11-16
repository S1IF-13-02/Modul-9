package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	hasil := "positif"
	if bilangan < 0 {
		hasil = "bukan positif"
	}
	fmt.Print(hasil)
}
