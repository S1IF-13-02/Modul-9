package main

import "fmt"

func main() {
	var a int

	fmt.Print("masukan nilai : ")
	fmt.Scan(&a)

	hasil := "bukan"
	if a < 0 && a%2 == 0 {
		hasil = "genap negatif"
	}
	fmt.Println(hasil)
}
