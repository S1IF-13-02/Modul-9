package main

import "fmt"

func main() {
	var a int
	var b = "bukan"
	fmt.Scan(&a)
	if a < 0 && a % 2 == 0 {
		b = "genap negatif"
	}
	fmt.Println(b)
}