package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("masukan angka:")
	fmt.Scan(&n)

	if n > 0{
		fmt.Println("positif")
	} else if n <= 0 {
		fmt.Println("bukan positif")
	}
}
