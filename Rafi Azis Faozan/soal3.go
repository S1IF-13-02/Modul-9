package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	Xfaktor := false
	if y%x == 0 {
		Xfaktor = true
	}
	fmt.Println("Nilai x bersifat: ", Xfaktor)

	Yfaktor := false
	if x%y == 0 {
		Yfaktor = true
	}
	fmt.Println("Nilai y bersifat: ", Yfaktor)
}
