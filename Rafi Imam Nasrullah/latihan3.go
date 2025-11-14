package main

import "fmt"

func main() {
	var nilai int
	var bilangan bool = false

	fmt.Print("masukan nilai: ")
	fmt.Scan(&nilai)

	if nilai < 0 && nilai%2 == 0 {
		bilangan = true
	}

	fmt.Println(bilangan)
}
