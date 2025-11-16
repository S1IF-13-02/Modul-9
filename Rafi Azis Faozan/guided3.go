package main

import "fmt"

func main() {
	var bilangan int
	var myBool bool
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	myBool = false
	if bilangan < 0 && bilangan%2 == 0 {
		myBool = true
	}
	fmt.Println(myBool)
}
