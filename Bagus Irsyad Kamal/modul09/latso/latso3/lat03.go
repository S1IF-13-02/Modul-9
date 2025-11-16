package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	hasil1 := y%x == 0
	hasil2 := x%y == 0

	fmt.Println(hasil1)
	fmt.Println(hasil2)
}
