package main

import "fmt"
func main() {
	var a,hasil int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&a)
	fmt.Println(a)
	if a < 0{
		hasil = a * -1
		fmt.Println(hasil)
	} 
}