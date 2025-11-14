package main

import "fmt"

func main() {
	var a, hasil int
	fmt.Print("masukan nilai : ")
	fmt.Scan(&a)

	if a%2 == 0 || a > 0 {
		hasil = a/2 + a%2
	}

	fmt.Println(hasil)
}
