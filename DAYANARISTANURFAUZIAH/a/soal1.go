package main

import "fmt"

func main() {
	var motor, orang int
	fmt.Print("jumlah orang :")
	fmt.Scan(&orang)
	motor = orang / 2

	if orang%2 != 0 {
		motor = motor + 1
	}
	fmt.Print("jumlah motor :", motor)
}
