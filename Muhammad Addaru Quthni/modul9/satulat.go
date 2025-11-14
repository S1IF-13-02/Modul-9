package main
import "fmt"
func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	if n < 0 {
		n = -n
	}
	fmt.Print("Nilai absolut: ", n)
}