package main 

import "fmt"

func main () {
	var x int
	var u bool = false

	fmt.Scan(&x) 

	if x < 0 && x%2 == 0  {
		u = true
	}
	fmt.Println(u)

}