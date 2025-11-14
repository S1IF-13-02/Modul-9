package main

import "fmt"

func main() {
	var a int
	var positiff = "positif"

	fmt.Print("masukan nilai : ")
	fmt.Scan(&a)

	if a <= 0 {
		positiff = "negatif"

	}
	fmt.Print(positiff)

}
