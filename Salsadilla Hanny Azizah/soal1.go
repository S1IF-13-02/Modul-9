package main

import "fmt"

func main() {
	var a,b int

	fmt.Print("Masukkan jumlah orang = ")
	fmt.Scan(&a)

	if a > 0 {
		b = (a / 2) + (a % 2)
	}
	fmt.Println("Jadi jumlah motor yang diperlukan untuk touring adalah", b)
}