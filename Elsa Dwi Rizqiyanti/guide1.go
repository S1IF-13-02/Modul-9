package main

import "fmt"

func main() {
	var absolut int

	fmt.Scan(&absolut)
	if absolut < 0 {
		absolut = -absolut
		fmt.Println("Hasil adalah", absolut)
	}
	fmt.Println(absolut)
}
