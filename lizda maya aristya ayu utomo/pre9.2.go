package main

import "fmt"

func main() {
	var n int
	var isPositif = "positif"

	fmt.Scan(&n)

	if n <= 0 {
		isPositif = "negatif"
	}
	fmt.Println(isPositif)

}
