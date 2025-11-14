package main
import "fmt"
func main() {
	var x, y int
	var xfaktory bool = true
	var yfaktorx bool = true

	fmt.Scan(&x, &y)
	if y%x != 0 {
		xfaktory = false
	}
	if x%y != 0 {
		yfaktorx = false
	}
	fmt.Println(xfaktory)
	fmt.Println(yfaktorx)
}