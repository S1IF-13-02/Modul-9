package main

import "fmt"

func main() {
	
	var x int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&x)

	if x < 0 && x % 2 == 0 {
		fmt.Print(x, " adalah bilangan genap negatif")
	}else{
		fmt.Print(x, " bukan bilangan genap negatif")
	}
}
