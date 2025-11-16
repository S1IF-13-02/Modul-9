package main

import (
	"fmt"
)

func main() {
	var a int
	var hasil int
	fmt.Print("masukan jumlah orang: ")
	fmt.Scan(&a)

	hasil = 0
	for i := 0; i < a; i+= 2 {
		hasil ++ 
	}
	fmt.Printf("motor: %d", hasil)
}
