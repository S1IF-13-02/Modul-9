package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Print("masukan 2 bilangan bulat: ")
	fmt.Scan(&A, &B)

	adalahFactorB := B % A ==0
	adalahFactorA := A % B ==0

	fmt.Println(A, adalahFactorB)
	fmt.Println(B, adalahFactorA)
}
