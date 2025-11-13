package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("bilangan (x y): ")
	fmt.Scan(&x, &y)

	xFaktorY := y % x == 0
	yFaktorX := x % y == 0

	fmt.Println(xFaktorY)
	fmt.Println(yFaktorX)
}
