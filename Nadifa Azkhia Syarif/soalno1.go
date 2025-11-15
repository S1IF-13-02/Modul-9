package main

import "fmt"

func main() {
	var orang int
	var motor int

	fmt.Scan(&orang)

	motor = orang / 2

	if orang%2 != 0 {
		motor = motor + 1
	}

	fmt.Println(motor)
}
