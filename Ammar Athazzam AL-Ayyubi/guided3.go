package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("masukan angka: ")
	fmt.Scan(&n)

	if n < 0 && n%2 == 0 {
		fmt.Printf("%d ---> true\n", n)
	} else {
		fmt.Printf("%d ---> false\n", n)
	}
}
