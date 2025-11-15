package main

import "fmt"
func main (){
	var x int
	i := "Bukan"
	fmt.Print("Masukan Bilangan bulat : ")
	fmt.Scan(&x)
	if x%2 == 0 && x<0 {
		i = "Genap Negatif"
	}
	fmt.Print(i)
}