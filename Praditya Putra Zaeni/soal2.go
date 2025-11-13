package main 

import "fmt"

func main () {
	var x int
	var u string = "bukan"

	fmt.Scan(&x) 

	if x <= 0 {
		u = "genap positif"
	}
	fmt.Println(u)

}