package main

import "fmt"

func main(){
var bilanganbulat int
fmt.Scan(&bilanganbulat)
if bilanganbulat < 0 {
	bilanganbulat = -1 * bilanganbulat 
	}
fmt.Println(bilanganbulat)
}