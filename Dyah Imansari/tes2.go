package main
import "fmt"
func main() {
	var x int
	var p string = "positif"
	
	fmt.Scan(&x)
	if x < 0 {
		p = "bukan positif"
	}
	fmt.Println(p)
}