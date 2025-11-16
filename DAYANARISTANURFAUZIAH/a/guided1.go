package main

import "fmt"

func main() {
	var a int
	fmt.Print("masukan angka: ")
	fmt.Scan(&a)

	if a < 0 {
		a = -a
	}
	fmt.Println(a)
}
