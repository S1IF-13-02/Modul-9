package main

import "fmt"
	
func main() {
	var x int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&x)

	if x < 0 {
		x = -x
	}

	fmt.Println("Output nya adalah:", x)
}
