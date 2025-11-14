package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("")
	fmt.Scan(&x)

	if x > 0 {
		fmt.Println("positif")
	} else {
		fmt.Println("bukan positif")
	}
}
