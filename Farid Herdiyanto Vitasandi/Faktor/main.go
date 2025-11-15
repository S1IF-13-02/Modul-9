package main

import "fmt"

func main() {
	
	var x, y int

	fmt.Print("Masukkan bilangan (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan (y): ")
	fmt.Scan(&y)

	output1 := false
	output2 := false

	if (y != 0 && x != 0) {
		output1 = y % x == 0
		output2 = x % y == 0
	}else{
		fmt.Println("Input tidak boleh nol")
		return	
	}
	fmt.Println(output1)
	fmt.Println(output2)
	
	
}
