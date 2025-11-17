package main

import "fmt"

func main() {
	var x, y int64
	var res1, res2 bool = true, true

	fmt.Print("Masukkan Angka (x, y): ")
	fmt.Scan(&x, &y)

	if y%x != 0 {
		res1 = false
	}

	if x%y != 0 {
		res2 = false
	}

	fmt.Println(res1, res2)
}
