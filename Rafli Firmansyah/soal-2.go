package main

import "fmt"

func main() {
	var a int
	fmt.Print("Masukan nilai a : ")
	fmt.Scan(&a)

	keterangan := "bukan"

	if a < 0 && a%2 == 0 {
		keterangan = "genap negatif "
	}

	fmt.Print(keterangan)
}
