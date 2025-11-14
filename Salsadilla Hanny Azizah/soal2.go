package main

import "fmt"

func main(){
	var bilangan int
	var isPositif = "bukan"

	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&bilangan)
	if bilangan < 0 {
		isPositif = "genap negatif"
	}
	fmt.Print(isPositif)
}