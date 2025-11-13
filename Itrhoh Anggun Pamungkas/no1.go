package main

import "fmt"

func main() {
	var orang, motor int
	fmt.Print("jumlah orang: ")
	fmt.Scan(&orang)

	motor = orang / 2 

	if orang%2 != 0 { 
		motor = motor + 1 
	}
	fmt.Println("motor yang diperlukan:", motor)
}
