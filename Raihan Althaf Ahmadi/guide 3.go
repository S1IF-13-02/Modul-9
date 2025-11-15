package main

import "fmt"

func main() {
	var a int
	hsl := false
	fmt.Print("Masukkan Bilangan Bulat: ")
	fmt.Scan(&a)
	if a < 0 && a%2 == 0 {
		hsl = true
	}
	fmt.Println(hsl)
}
