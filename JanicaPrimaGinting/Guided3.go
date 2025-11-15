package main 

import "fmt"

func main(){
	var bilanganbulat int 
	var hasil bool 
	fmt.Scan(&bilanganbulat)
	if bilanganbulat < 0 && bilanganbulat % 2 == 0 {
		hasil = true
	}
	fmt.Print(hasil)
}