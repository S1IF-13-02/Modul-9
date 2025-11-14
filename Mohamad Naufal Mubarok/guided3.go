package main

import "fmt"

func main() {
	var n int
	var pinter bool = false

	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&n)

	if n < 0 && n%2 == 0 {
		pinter = true
		
	}
	fmt.Println(pinter)
}