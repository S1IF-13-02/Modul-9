package main

import "fmt"

func main() {
	var x int64
	var tmp bool = false

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&x)

	if x < 0 && x%2 == 0 {
		tmp = true
	}

	fmt.Println(tmp)
}
