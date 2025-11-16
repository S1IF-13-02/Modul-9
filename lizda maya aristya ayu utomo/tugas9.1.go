package main

import "fmt"

func main() {
	var jumlahorang, jumlahmotor int

	fmt.Print("masukan jumlah orang: ")
	fmt.Scan(&jumlahorang)

	jumlahmotor = jumlahorang / 2
	if jumlahorang%2 != 0 {
		jumlahmotor++
	}
	fmt.Println("jumlah motor yang diperlukan: ", jumlahmotor)
}
