package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	// x adalah faktor dari y jika y habis dibagi x
	fmt.Println(y%x == 0)

	// y adalah faktor dari x jika x habis dibagi y
	fmt.Println(x%y == 0)
}
