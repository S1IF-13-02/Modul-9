package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var xFaktorY bool = y%x == 0

	var yFaktorX bool = x%y == 0

	fmt.Println(xFaktorY)
	fmt.Println(yFaktorX)
}
