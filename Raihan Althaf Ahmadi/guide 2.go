package main

import "fmt"
func main (){
	var a int
	isPositif := "Positif"
	fmt.Print("Masukan Bilangan Bulat : ")
	fmt.Scan(&a)
	if a <= 0 {
		isPositif = "Tidak Positif"
	}
	fmt.Println(isPositif)
}