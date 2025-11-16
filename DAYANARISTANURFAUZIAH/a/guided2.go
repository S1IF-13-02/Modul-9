package main

import "fmt"

func main() {
	var a int
	var b string
	fmt.Print("masukan bilangan: ")
	fmt.Scan(&a)
	b = "bukan positif"

	if a > 0 {

		b = "positif"

	}
	fmt.Println(b)

}
