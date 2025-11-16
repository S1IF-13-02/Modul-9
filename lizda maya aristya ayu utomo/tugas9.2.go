package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("masukan bilangan: ")
	fmt.Scan(&bilangan)

	status := "bukan"

	if bilangan < 0 && bilangan%2 == 0 {
		status = "genap negatif"
	}
	fmt.Println(status)

}
