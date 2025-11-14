package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan dua bilangan: ")
	fmt.Scan(&x, &y)

	if true {
		fmt.Println(y%x == 0)
		fmt.Println(x%y == 0)
	}
}
