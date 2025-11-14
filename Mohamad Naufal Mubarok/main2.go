package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 && bilangan < 0 {
		fmt.Println("genap negatif")
		return
	}
	fmt.Println("Jawabannya adalah: Bukan")
}