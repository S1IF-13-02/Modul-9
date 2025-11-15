package main

import "fmt"
func main () {
	var x, y int
	fmt.Print("Masukan dua bilangan bulat : ")
	fmt.Scan(&x, &y)
	faktorx := false
	faktory := false
	
	if y%x==0{
		faktorx = true
	}


	if x%y==0{
		faktory = true
	}

	fmt.Println(faktorx)
	fmt.Println(faktory)
}