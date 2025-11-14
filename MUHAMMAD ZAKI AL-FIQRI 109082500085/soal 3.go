package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukan x : ")
	fmt.Scan(&x)
	fmt.Print("masukan y : ")
	fmt.Scan(&y)
	if true {
		fmt.Println(y%x == 0)
		fmt.Println(x%y == 0)
	}
}
