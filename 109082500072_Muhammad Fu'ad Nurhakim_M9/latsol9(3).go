package main

import "fmt"

func main() {
	var x, y int
	var a, b bool

	fmt.Scan(&x, &y)
	if x/y == 0 || x/y == 1 {
		a = true
	}
	if y/x == 0 || y/x == 1 {
		b = true
	}
	fmt.Println(a)
	fmt.Println(b)
}
