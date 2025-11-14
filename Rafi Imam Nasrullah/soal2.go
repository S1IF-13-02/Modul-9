package main

import "fmt"

func main(){
	var nilai int
	var p string = "bukan"

	fmt.Print("masukan nilai: ")
	fmt.Scan(&nilai)	

	if nilai < 0 {
		p = "genap negatif"
	}
	fmt.Println(p)
}