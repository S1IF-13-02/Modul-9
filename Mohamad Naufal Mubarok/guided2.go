package main

import "fmt"

func main() {
	var x float64
	var yaya string = "positif"

	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&x)

	if x <= 0 {
		yaya = "bukan positif"

	}
	fmt.Println(yaya)
}
