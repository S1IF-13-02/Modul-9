package main

import "fmt"

func main() {
	var bil int
	var hasil = "false"

	fmt.Scan(&bil)

	if bil < 0 && bil%2 == 0 {
		hasil = "true"
	}

	fmt.Println(hasil)
}
