package main

import "fmt"

func main() {
	var bil int
	var isPositif = "positif"

	fmt.Scan(&bil)
	if bil <= 0  {
		isPositif = "negatif"
	}
	fmt.Println(isPositif)
}