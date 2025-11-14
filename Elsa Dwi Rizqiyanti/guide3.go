package main

import "fmt"

var isPositif bool = false
var n int

func main() {

	fmt.Scan(&n)

	if n < 0 && n%2 == 0 {
		isPositif = true
	}

	fmt.Println(isPositif)
}
