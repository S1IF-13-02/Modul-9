package main 

import "fmt"

func main () {
	var x int
	var u string = "positif"

	fmt.Scan(&x) 

	if x <= 0 {
		u = "bukanpositif"

	}
	fmt.Println(u)

}