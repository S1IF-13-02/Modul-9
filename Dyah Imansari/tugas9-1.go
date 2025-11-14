package main
import "fmt"
func main() {
	var m, n int

	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scanln(&n)
	m = n/2
	if n%2 != 0 {
		m = m + 1
	}
	fmt.Print("Jumlah motor adalah ", m)
}