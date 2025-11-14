package main

import "fmt"

func main() {
	var x, y int
	var faktor1, faktor2 bool

	fmt.Print("Masukkan bilangan (x): ")
	fmt.Scan(&x)

	fmt.Print("Masukan bilangan (y)")
	fmt.Scan(&y)

	if y%x == 0 {
		faktor1 = true
	}
	if y%x != 0 {
		faktor1 = false
	}

	if x%y == 0 {
		faktor2 = true
	}
	if x%y != 0 {
		faktor2 = false
	}

	fmt.Println(faktor1)
	fmt.Println(faktor2)
}