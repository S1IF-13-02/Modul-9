package main

import "fmt"

func main() {
	var bilangan int
	var teks string
	fmt.Scan(&bilangan)
	teks = "Bukan genap negatif"
	if bilangan < 0 {
		teks = "Genap negatif"
	}
	fmt.Println(teks)
}
