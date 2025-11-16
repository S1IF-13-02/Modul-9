package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	var hasil string = "bukan"

	if bilangan%2 == 0 && bilangan < 0 {
		hasil = "genap negatif"
	}

	fmt.Println(hasil)
}
