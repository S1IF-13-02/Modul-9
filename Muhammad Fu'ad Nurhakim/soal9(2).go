package main

import "fmt"

func main() {
	var a int
	var b string = "positif"

	fmt.Scan(&a)
	if a < 0 {
		b = "bukan positif"
	}
	fmt.Print(b)
}
