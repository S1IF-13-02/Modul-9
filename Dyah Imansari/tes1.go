package main
import "fmt"
func main() {
	var x int
	fmt.Scan(&x)
	if x < 0 {
		x = -x
	}
	fmt.Println("Nilai mutlak adalah", x)
	fmt.Println("Selesai")
}