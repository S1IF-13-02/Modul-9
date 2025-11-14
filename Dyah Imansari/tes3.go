package main
import "fmt"
func main() {
	var b int
	var p bool = false

	fmt.Scan(&b)
	if b<0 && b%2 == 0 {
		p = true
	}
	fmt.Println(p)
}