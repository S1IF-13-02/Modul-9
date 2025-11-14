package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan nilai n : ")
	fmt.Scan(&n)

	if n < 0 {
		n = -n
	}
	fmt.Print("Nilai Absolut/mutlak : ", n)
}
