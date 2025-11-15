package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Nilai y: ")
	fmt.Scan(&y)
	fmt.Println(y%x == 0)
	fmt.Println(x%y == 0)
}
