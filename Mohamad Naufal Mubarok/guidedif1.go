package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&nilai)

	if nilai < 0 {
	nilai = -nilai
	}
	fmt.Println(nilai)
}	