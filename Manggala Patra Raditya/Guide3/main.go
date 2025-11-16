package main

import "fmt"

func main(){
	var bilangan int
	var isPositif bool

	fmt.Scan(&bilangan)
	if bilangan < 0 && bilangan % 2 == 0{
		isPositif = true
	}
	fmt.Print(isPositif)
}