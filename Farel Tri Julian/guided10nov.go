package main

import "fmt"

func main() {
	var a int
	fmt.Print("masukan nilai : ")
	fmt.Scan(&a)

	if a < 0 {
		a = -a
	}

	fmt.Println(a)
}
