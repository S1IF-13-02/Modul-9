package main

import "fmt"

func main() {
	var a int
	var b bool = false

	fmt.Scan(&a)
	if a < 0 && a%2 == 0 {
		b = true
	}
	fmt.Print(b)
}
