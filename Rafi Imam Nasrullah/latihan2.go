package main

import "fmt"

func main(){
	var nilai int
	var p string = "positif"

	fmt.Print("masukan nilai: ")
	fmt.Scan(&nilai)	

	if nilai <= 0 {
		p = "bukan positif"
	}
	fmt.Println(p)
}