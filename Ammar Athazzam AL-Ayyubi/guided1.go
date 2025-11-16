package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("masukan angka: ")
	fmt.Scan(&n)

	if n < 0 {
		n = -n
	}
	fmt.Println(n)
}
