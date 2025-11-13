package main 

import "fmt"

func main () {
	var x, y int

	fmt.Scan(&x,&y) 

	if true{fmt.Println(y%x == 0)}
	if true{fmt.Println(x%y == 0)}
}