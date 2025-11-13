package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan dua bilangan (x y): ")
	fmt.Scan(&x, &y)

	fmt.Println(y%x == 0)
	fmt.Println(x%y == 0)
}
