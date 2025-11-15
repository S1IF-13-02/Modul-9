package main

import "fmt"

func main() {
	
	var bilangan int
	var teks string

	fmt.Scan(&bilangan)

	teks = "Bukan positif"

	if bilangan > 0{
		teks = "Positif"
	}
	fmt.Println(teks)
}
