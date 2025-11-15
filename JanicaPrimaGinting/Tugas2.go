package main

import "fmt"

func main(){
	var bilanganbulat int
	var posisi = "bukan"
	fmt.Scan(&bilanganbulat)
	if bilanganbulat < 0 && bilanganbulat % 2 == 0{
		posisi = "genap positif" 
	}
	fmt.Println(posisi)
}