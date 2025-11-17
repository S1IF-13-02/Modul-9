package main

import "fmt"

func main() {
	var x int64
	var tmp string = "positif"

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&x)

	if x < 0 {
		tmp = "bukan positif"
	}

	fmt.Println(tmp)
}
