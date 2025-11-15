package main

import "fmt"

func main(){
	var bilanganbulat int
	var hasil string = "bukan positif"
	fmt.Scan(&bilanganbulat)
	if bilanganbulat > 0 { 
		hasil = "positif"
	}
	fmt.Print(hasil)
}