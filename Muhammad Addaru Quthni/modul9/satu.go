package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Masukkan Jumlah Orang : ")
	fmt.Scan(&a)
	if a % 2 == 0{
		b = a / 2
	}
	b = (a / 2) + (a % 2)
	fmt.Println("Jumlah Motor untuk Touring adalah :",b)
}