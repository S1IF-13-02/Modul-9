package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("masukan dua bilangan :")
	fmt.Scan(&x, &y)
	xfaktory := y%x == 0
	yfaktorx := x%y == 0

	fmt.Println(xfaktory)
	fmt.Println(yfaktorx)
}
