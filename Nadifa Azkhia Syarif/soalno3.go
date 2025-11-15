package main

import "fmt"

func main() {
	var x, y int
	var xFaktorY bool
	var yFaktorX bool

	fmt.Scan(&x, &y)

	xFaktorY = (y%x == 0)

	yFaktorX = (x%y == 0)

	fmt.Println(xFaktorY)
	fmt.Println(yFaktorX)
}
